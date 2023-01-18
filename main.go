/*
Copyright © 2022 javiroman <javiroman@apache.org>

*/
package main

import (
	"flag"
	"github.com/javiroman/propgen/cmd"
	"os"
)

var (
	label  *string
	render *string
	file   *string
)

var Version = "development"

func init() {
	label = flag.String("label", "", "NIFI|HADOOP|ZOOKEEPER|KUDU|HIVE|IMPALA")
	render = flag.String("render", "", "coresite|hdfssite|nifiproperties|zoocfg|metastoresite|hivesite")
	file = flag.String("file", "", "/path/to/filename")
	flag.Parse()

	if flag.NFlag() < 3 {
		usage()
		os.Exit(1)
	}
	checkParametersIntegrity()
}

func main() {
	cmd.GenerateFile(*label, *render, *file)
}
