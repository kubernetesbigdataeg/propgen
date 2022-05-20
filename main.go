/*
Copyright © 2022 javiroman <javiroman@apache.org>

*/
package main

import (
	"flag"
	"fmt"
	"github.com/javiroman/propgen/cmd"
	"os"
)

var (
	label *string
	file  *string
)

func init() {
	label = flag.String("label", "", "NIFI|HADOOP|ZOOKEEPER|KUDU")
	file = flag.String("file", "", "/path/to/filename")
	flag.Parse()

	if len(os.Args[1:]) <= 2 {
		fmt.Println("Usage: propgen [flags] \nFlags:")
		flag.PrintDefaults()
		fmt.Println("Examples:")
		fmt.Println("\tpropgen -label NIFI -file /opt/nifi/conf/nifi.properties")
		fmt.Println("\tpropgen -label HADOOP -file /opt/hadoop/etc/hadoop/conf/core-site.xml")
		os.Exit(1)
	}
}

func main() {
	cmd.GenerateFile(*label, *file)
}
