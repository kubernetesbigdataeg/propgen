/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const (
	RENDER int = 1
	VALUES     = 2
)

func SaveReplace(cad string) string {
	return strings.Replace(strings.SplitN(cad, "=", 2)[0], "_", ".", -1)
}

func SaveSplit(cad string) string {
	return strings.SplitN(cad, "=", 2)[1]
}

func GenerateFile(label string, render string, file string) {
	configEntries := make(map[string]string)
	for _, e := range os.Environ() {
		matchedLabel, _ := regexp.MatchString(fmt.Sprintf("^%s", label), e)
		subcad := strings.SplitN(e, "__", 3)
		if matchedLabel && (subcad[RENDER] == render) {
			configEntries[SaveReplace(subcad[VALUES])] = SaveSplit(subcad[VALUES])
		}
	}

	switch render {
	case ZooCfg:
		fmt.Println("Rendering Zookeeper config file ...")
		tmp1, _ := template.New(ZooCfg).Parse(zoocfg)
		err := tmp1.Execute(os.Stdout, configEntries)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(file, ":")
	case CoreSite, HdfsSite:
		fmt.Println("Rendering Hadoop config file ...")
		fmt.Println("Rendering Zookeeper config file ...")
		tmp1, _ := template.New(CoreSite).Parse(coresite)
		err := tmp1.Execute(os.Stdout, configEntries)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(file, ":")
	case NifiProperties:
		fmt.Println("Rendering Hadoop config file ...")
	default:
		fmt.Println("No render found")
	}
}
