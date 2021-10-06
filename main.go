package main

import (
	"log"
	"os"
	"strings"

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
	dup := make(map[int]string)

check:
	for _, c := range checksList {
		for _, j := range unique {
			if c.Name == j {
				dup[c.ID] = c.Name
				continue check
			}
		}
		unique = append(unique, c.Name)
	}
	if len(dup) > 0 {
		log.Printf("Found %v duplicated checks: %v\n", len(dup), dup)
	}
	log.Println(len(unique), "Unique checks")

	removeDups := os.Getenv("REMOVE_DUPLICATED")
	if strings.ToLower(removeDups) == "true" {
		log.Println("Start to remove duplicated checks:")
		for id, name := range dup {
			log.Printf("Removing %v with ID: %v\n", name, id)
			removeCheck, err := client.Checks.Delete(id)
			check(err)
			log.Println(removeCheck)
		}
	}
}
