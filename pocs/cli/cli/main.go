/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ltmleo/golang/pocs/cli/cli/LojaOnline/iac/iac-lojaonline-envs/cli/private"
	"gopkg.in/yaml.v3"
)

func main() {
	var config map[string]map[string]string
	yamlFile, _ := ioutil.ReadFile("/home/vivo/Documents/VIVO/iac-lojaonline-envs/agreement.yaml")
	_ = yaml.Unmarshal(yamlFile, &config)
	fmt.Println(config)
	private.Execute()

}
