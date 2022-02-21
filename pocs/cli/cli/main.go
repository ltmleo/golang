/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"10.129.178.173/LojaOnline/iac/iac-lojaonline-envs/cli/private"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"fmt"
	)

func main() {
	var config map[string]map[string]string
	yamlFile, _ := ioutil.ReadFile("/home/vivo/Documents/VIVO/iac-lojaonline-envs/agreement.yaml")
	_ = yaml.Unmarshal(yamlFile, &config)
	fmt.Println(config)
	private.Execute()

}
