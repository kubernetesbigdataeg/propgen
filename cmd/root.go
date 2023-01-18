/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
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
	if !strings.Contains(cad, "=") {
		log.Fatal("Unable get Key (left hand =) in: ")
	}
	strReplaced := strings.Replace(strings.SplitN(cad, "=", 2)[0], "___", "-", -1)
	return strings.Replace(strings.SplitN(strReplaced, "=", 2)[0], "_", ".", -1)
}

// SafeGetKey Gets the key without change anything
func SafeGetKey(cad string) string {
	if !strings.Contains(cad, "=") {
		log.Fatal("Unable get Key (left hand =) in: ")
	}
	return strings.SplitN(cad, "=", 2)[0]
}

// SafeGetValue Gets the value
func SafeGetValue(cad string) string {
	if !strings.Contains(cad, "=") {
		log.Fatal("Unable get Value (right hand =) in: ")
	}
	return strings.SplitN(cad, "=", 2)[1]
}

func GenerateFile(label string, render string, file string) {
	configEntries := make(map[string]string)
	var subcad []string

	for _, e := range os.Environ() {
		matchedLabel, _ := regexp.MatchString(fmt.Sprintf("^%s", label), e)

		// If variable begins with "label" but not well formed print warning.
		if matchedLabel && strings.Count(e, "__") < 2 {
			log.Println("WARNING: Not found separator in -> ", e)
		} else {
			subcad = strings.SplitN(e, "__", 3)
			if matchedLabel && (subcad[RENDER] == render) {
				if label == "KUDU" || (label == "IMPALA" && 
				  (render == "impaladaemon" || render == "impalacatalog" || render == "impalastatestore" || render == "impalaadmission")) {
					configEntries[SafeGetKey(subcad[VALUES])] = SafeGetValue(subcad[VALUES])
				} else {
					configEntries[SafeGetKeyReplaceUnderscoreByDot(subcad[VALUES])] =
						SafeGetValue(subcad[VALUES])
				}
			}
		}
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	switch render {
	case ZooCfg:
		tmp1, _ := template.New(ZooCfg).Parse(zoocfg)
		err := tmp1.Execute(f, configEntries)
		if err != nil {
			fmt.Println(err)
		}
	case CoreSite, HdfsSite:
		tmp1, _ := template.New(CoreSite).Parse(coresite)
		err := tmp1.Execute(f, configEntries)
		if err != nil {
			fmt.Println(err)
		}
	case NifiProperties:
		tmp1, _ := template.New(NifiProperties).Parse(nifiproperties)
		err := tmp1.Execute(f, configEntries)
		if err != nil {
			fmt.Println(err)
		}
	case KuduMaster, KuduTserver, ImpalaDaemon, ImpalaCatalog, ImpalaStateStore, ImpalaAdmission:
		tmp1, _ := template.New(KuduMaster).Parse(kudumaster)
		err := tmp1.Execute(f, configEntries)
		if err != nil {
			fmt.Println(err)
		}
	case MetastoreSite, HiveSite:
		tmp1, _ := template.New(MetastoreSite).Parse(metastoresite)
		err := tmp1.Execute(f, configEntries)
		if err != nil {
			fmt.Println(err)
		}
	default:
		log.Fatal("No render match found!")
	}

}
