package main

import (
	"log"
	"os"

	"github.com/russellcardullo/go-pingdom/pingdom"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	client, err := pingdom.NewClientWithConfig(pingdom.ClientConfig{
		APIToken: os.Getenv("PINGDOM_API_TOKEN"),
	})
	check(err)

	checksList, err := client.Checks.List()
	check(err)

	log.Println(len(checksList), "Total checks")

	unique := make([]string, 0)
	dup := make(map[string]int)
check:
	for _, c := range checksList {
		for _, j := range unique {
			if c.Name == j {
				dup[c.Name] += 1
				continue check
			}
		}
		unique = append(unique, c.Name)
	}
	if len(dup) > 0 {
		log.Println("Found duplicated checks: ", dup)
	}
	log.Println(len(unique), "Unique checks")
}
