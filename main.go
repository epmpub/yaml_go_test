package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Person struct {
	Name    string   `yaml:"name"`
	Age     int      `yaml:"age"`
	Hobbies []string `yaml:"hobbies"`
}

func main() {
	Person := Person{
		Name:    "Kevin",
		Age:     23,
		Hobbies: []string{"Cpp", "Reading", "Writing", "Coding", "Funning"},
	}

	yamlFile, err := yaml.Marshal(&Person)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(yamlFile))

	f, err := os.Create("person.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.WriteString(f, string(yamlFile))
	if err != nil {
		panic(err)
	}

	Read()
	fmt.Println("------------------------------------------")
	Read2()
	fmt.Println("------------------------------------------")
	Read3()

}
