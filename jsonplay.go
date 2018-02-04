package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Projects struct {
	Projects []struct {
		Name     string `json:"name"`
		Services []struct {
			Name      string   `json:"name"`
			Instances []string `json:"instances"`
		} `json:"services"`
	} `json:"projects"`
}

func (p *Projects) GetInfo() string {
	r := "List of projects: \n"
	for n, project := range p.Projects {
		r += fmt.Sprintf("\nProject %d: %s\n", n, project.Name)
		r += fmt.Sprintf("\tServices:\n")
		for sn, srv := range project.Services {
			r += fmt.Sprintf("\t\tservice %d: %s\n", sn, srv.Name)
			for ki, inst := range srv.Instances {
				r += fmt.Sprintf("\t\t\tinstance %d: %s\n", ki, inst)
			}
		}
	}
	return r
}

func main() {
	fcontext, err := ioutil.ReadFile("./my.json")
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	fmt.Println(string(fcontext))
	projectContent := &Projects{}
	json.Unmarshal(fcontext, projectContent)
	fmt.Println(projectContent.GetInfo())

}

