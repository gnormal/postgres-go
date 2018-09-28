package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"gnorm.org/gnorm/run/data"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Printf("usage: %v <function>", os.Args[0])
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Printf("error reading from stdin: %v", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "insertCols":
		if err := insertNames(b, dbName); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	case "insertFields":
		if err := insertNames(b, fieldName); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	default:
		log.Printf("unknown function: %v", os.Args[1])
		os.Exit(1)
	}
}

func insertNames(b []byte, f func(*data.Column) string) error {
	data := struct {
		Cols data.Columns `json:"data"`
	}{}
	if err := json.Unmarshal(b, &data); err != nil {
		return fmt.Errorf("can't unmarshal json input: %v", err)
	}
	// this needs to be non-nil or json marshalling gets confused when it's empty
	names := []string{}
	for _, c := range data.Cols {
		// don't insert these since they're auto-updated
		if c.DBName == "updated_at" || c.DBName == "created_at" {
			continue
		}
		// don't insert generated IDs
		if c.IsPrimaryKey && c.HasDefault {
			continue
		}
		names = append(names, f(c))
	}
	sort.Strings(names)
	b, err := json.Marshal(map[string]interface{}{
		"data": names,
	})
	if err != nil {
		return fmt.Errorf("error translating output to json: %v", err)
	}
	if _, err := os.Stdout.Write(b); err != nil {
		return fmt.Errorf("error writing to stdout: %v", err)
	}
	return nil
}

func dbName(c *data.Column) string {
	return c.DBName
}

func fieldName(c *data.Column) string {
	return "r." + c.Name
}
