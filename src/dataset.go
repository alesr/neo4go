package neo4go

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/spf13/viper"
)

const (
	connectErrMsg    = "Error connectiong to Neo4J"
	createNodeErrMsg = "Error creating node"
	prepStmtErrMsg   = "Error preparing statement"
	execStmtErrMsg   = "Error running statement"
)

var conn bolt.Conn

// Load data set to cluster
func Load() {

	// Load Viper configuration
	loadConfig()

	// Initialize Auth struct and get URL for Neo4j connection
	a := newAuth()
	a.getURL()

	// Ask for a new Golang Neo4J Bolt Driver
	driver := bolt.NewDriver()

	// Open new connection with Neo4j
	conn, err := driver.OpenNeo(a.URL)
	if err != nil {
		log.Fatalf("%s: %s", connectErrMsg, err)
	}
	defer conn.Close()

	// Load CSV data to an array containing queries for graph creation
	queries, err := assembly()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare Pipeline with queries to execute
	pipeline, err := conn.PreparePipeline(queries...)
	if err != nil {
		log.Fatal(err)
	}

	// Execute queries
	_, err = pipeline.ExecPipeline(nilSlice(queries)...)
	if err != nil {
		log.Fatal(err)
	}

	// Close pipeline
	if err = pipeline.Close(); err != nil {
		log.Fatal(err)
	}
}

// Read CSV dataset and create queries
func assembly() ([]string, error) {
	file, err := os.Open(viper.GetString("dataset.file1"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var queries []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		char := escapeQuote(row[0])
		queries = append(queries, addCharQuery(char))

		if row[1] != "None" {
			house := escapeQuote(row[1])
			queries = append(queries, addHouseQuery(house), addIsAllyQuery(char, house))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return queries, nil
}

// Assembly a query to create a new character
func addCharQuery(name string) string {
	return fmt.Sprintf("MERGE (c:Character{name: \"%s\"})", name)
}

// Assembly a query to create a new house
func addHouseQuery(house string) string {
	return fmt.Sprintf("MERGE (h:House{name: \"%s\"})", house)
}

// Assembly a query to create a new allegiance
func addIsAllyQuery(char, house string) string {
	return fmt.Sprintf("MATCH (c:Character{name: '%s'}), (h:House{name: '%s'}) MERGE (c)-[:HAS_ALLIANCE_WITH]->(h)", char, house)
}

// Escape quotes
func escapeQuote(s string) string {
	index := strings.Index(s, "'")
	if index != -1 {
		return s[:index] + "\\" + s[index:]
	}
	return s
}

// Return a slice of nil values with queries length.
// Workaround for creating n nil params for ExecPipeline method
func nilSlice(s []string) []map[string]interface{} {
	tmp := make([]map[string]interface{}, len(s), len(s))
	for i := range tmp {
		tmp[i] = nil
	}
	return tmp
}
