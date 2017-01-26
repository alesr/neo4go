package neo4go

import (
	"fmt"

	"github.com/spf13/viper"
)

// Auth ...
type Auth struct {
	user     string
	password string
	baseURL  string
	port     string
	URL      string
}

// newAuth returns an auth instance
func newAuth() *Auth {
	return &Auth{
		user:     viper.GetString("neo4j.user"),
		password: viper.GetString("neo4j.password"),
		baseURL:  viper.GetString("neo4j.localhost"),
		port:     viper.GetString("neo4j.port"),
	}
}

// getURL creates a connection URL
func (a *Auth) getURL() {
	a.URL = fmt.Sprintf("bolt://%s:%s@%s:%s/db/data",
		a.user, a.password, a.baseURL, a.port)
}
