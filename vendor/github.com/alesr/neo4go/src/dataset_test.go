package neo4go

import "testing"

func TestAddCharQuery(t *testing.T) {
	expected := "MERGE (c:Character{name: 'tom hanks'})"
	observed := addCharQuery("tom hanks")
	if expected != observed {
		t.Errorf("Expected %s, got %s", expected, observed)
	}
}

func TestAddHouseQuery(t *testing.T) {
	expected := "MERGE (h:House{name: 'alabama'})"
	observed := addHouseQuery("alabama")
	if expected != observed {
		t.Errorf("Expected %s, got %s", expected, observed)
	}
}

func TestAddIsAllyQuery(t *testing.T) {
	expected := "MATCH (c:Character{name: 'tom hanks'}), (h:House{name: 'alabama'}) MERGE (c)-[:HAS_ALLIANCE_WITH]->(h)"
	observed := addIsAllyQuery("tom hanks", "alabama")
	if expected != observed {
		t.Errorf("Expected %s, got %s", expected, observed)
	}
}

func TestNilSlice(t *testing.T) {
	dummySlice := make([]string, 10, 10)
	expected := len(dummySlice)
	observed := nilSlice(dummySlice)
	if expected != len(observed) {
		t.Errorf("Expected %d, got %d", expected, len(observed))
	}
}

var escapeQuoteTestCases = []struct {
	input, expected string
}{
	{"Foo", "Foo"},
	{"Bar", "Bar"},
	{"Foo'bar", "Foo\\'bar"},
}

func TestEscapeQuote(t *testing.T) {
	for _, test := range escapeQuoteTestCases {
		observed := escapeQuote(test.input)
		if observed != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, observed)
		}
	}
}
