package neo4go

import (
	"reflect"
	"testing"
)

func TestNewAuth(t *testing.T) {
	expected := new(Auth)
	observed := newAuth()
	if !reflect.DeepEqual(expected, observed) {
		t.Errorf("Expected %T, got %T", expected, observed)
	}
}

func TestGetURL(t *testing.T) {
	a := newAuth()
	a.getURL()
	expected := "bolt://:@:/db/data"
	// URL without values
	if a.URL != expected {
		t.Errorf("Expected %s, got %s", expected, a.URL)
	}
}
