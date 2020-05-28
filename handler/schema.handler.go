package handler

import (
	"fmt"
	"io/ioutil"
)

// GetSchema takes care of returning the schema for Grahpql
func GetSchema() (string, error) {
	schemaFileString, err := getSchema("./schema.graphql")
	if err != nil {
		fmt.Printf("There was an error reading the schema, %s", err)
		return "", err
	}

	return string(schemaFileString), nil
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
