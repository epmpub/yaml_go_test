package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Blog struct {
	Contributors []string `yaml:"contributors"`
	Title        string   `yaml:"title"`
	Date         string   `yaml:"date"`
	Tags         []string `yaml:"tags"`
	Draft        bool     `yaml:"draft"`
}

func Read() {
	var blog Blog
	yamlFile, err := os.ReadFile("blog.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &blog)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Contributors: %v \n %v \n", blog.Contributors[0], blog.Contributors[1])
	fmt.Printf("Title: %v \n", blog.Title)
	fmt.Printf("Date: %v \n", blog.Date)
	fmt.Printf("Tags: %v \n %v \n %v \n %v \n", blog.Tags[0], blog.Tags[1], blog.Tags[2], blog.Tags[3])
	fmt.Printf("Draft: %v \n", blog.Draft)

}

func Read2() {
	var blog map[string]interface{}

	yamlFile, err := os.ReadFile("blog.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &blog)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Contributor: %v \n %v \n", blog["contributors"].([]interface{})[0], blog["contributors"].([]interface{})[1])
	fmt.Printf("Title: %v \n", blog["title"])
	fmt.Printf("Date: %v \n", blog["date"])
	fmt.Printf("Tags: %v \n %v \n", blog["tags"].([]interface{})[0], blog["tags"].([]interface{})[1])
	fmt.Printf("Draft: %v \n", blog["draft"])

}

func Read3() {
	var config map[string]interface{}

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Language: %v \n", config["runtime"])
}
