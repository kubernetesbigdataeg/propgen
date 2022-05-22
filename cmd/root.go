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
	TEMPLATE int = 1
	VALUES       = 2
)

func SaveReplace(cad string) string {

}

func GenerateFile(label string, file string) {
	p := make(map[string]string)
	for _, e := range os.Environ() {
		matched, _ := regexp.MatchString(fmt.Sprintf("^%s", label), e)
		if matched {
			subcad := strings.SplitN(e, "__", 3)
			//fmt.Print(subcad[TEMPLATE])
			//fmt.Println(subcad[VALUES])
			p[strings.Replace(strings.SplitN(subcad[VALUES], "=", 2)[0], "_", ".", -1)] =
				strings.SplitN(subcad[VALUES], "=", 2)[1]
		}
	}

	fmt.Println(file, ":")

	tmp1, _ := template.New("zoocfg").Parse(zoocfg)
	err := tmp1.Execute(os.Stdout, p)

	if err != nil {
		fmt.Println(err)
	}
}
