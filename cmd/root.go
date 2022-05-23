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

// SafeGetKeyReplaceUnderscoreByDot Gets the key and change underscore by dots.
func SafeGetKeyReplaceUnderscoreByDot(cad string) string {
	return strings.Replace(strings.SplitN(cad, "=", 2)[0], "_", ".", -1)
}

// SafeGetKey Gets the key without change anything
func SafeGetKey(cad string) string {
	return strings.SplitN(cad, "=", 2)[0]
}

// SafeGetValue Gets the value
func SafeGetValue(cad string) string {
	return strings.SplitN(cad, "=", 2)[1]
}

func GenerateFile(label string, render string, file string) {
	configEntries := make(map[string]string)
	for _, e := range os.Environ() {
		matchedLabel, _ := regexp.MatchString(fmt.Sprintf("^%s", label), e)
		subcad := strings.SplitN(e, "__", 3)
		if matchedLabel && (subcad[RENDER] == render) {
			if label == "KUDU" || label == "IMPALA" {
				configEntries[SafeGetKey(subcad[VALUES])] = SafeGetValue(subcad[VALUES])
			} else {
				configEntries[SafeGetKeyReplaceUnderscoreByDot(subcad[VALUES])] =
					SafeGetValue(subcad[VALUES])
			}
		}
	}

	fmt.Println("To File: ", file)

	switch render {
	case ZooCfg:
		tmp1, _ := template.New(ZooCfg).Parse(zoocfg)
		err := tmp1.Execute(os.Stdout, configEntries)
		if err != nil {
			fmt.Println(err)
		}
	case CoreSite, HdfsSite:
		tmp1, _ := template.New(CoreSite).Parse(coresite)
		err := tmp1.Execute(os.Stdout, configEntries)
		if err != nil {
			fmt.Println(err)
		}
	case NifiProperties:
		tmp1, _ := template.New(NifiProperties).Parse(nifiproperties)
		err := tmp1.Execute(os.Stdout, configEntries)
		if err != nil {
			fmt.Println(err)
		}
	case KuduMaster, KuduWorker, ImpalaDaemon, ImpalaCatalog, ImpalaStateStore, ImpalaAdmission:
		tmp1, _ := template.New(KuduMaster).Parse(kudumaster)
		err := tmp1.Execute(os.Stdout, configEntries)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("No render found")
	}
}
