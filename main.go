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
	dupItem := make(map[string]int)
	dupList := make([]map[string]int, 0)
check:
	for _, c := range checksList {
		for _, j := range unique {
			if c.Name == j {
				dupItem[c.Name] += 1
				dupList = append(dupList, dupItem)
				continue check
			}
		}
		unique = append(unique, c.Name)
	}
	if len(dupList) > 0 {
		log.Println("Found duplicated checks: ", dupList)
	}
	log.Println(len(unique), "Unique checks")
}
