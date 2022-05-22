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
	label  *string
	render *string
	file   *string
)

func init() {
	label = flag.String("label", "", "NIFI|HADOOP|ZOOKEEPER|KUDU")
	render = flag.String("render", "", "coresite|hdfssite|nifiproperties|zoocfg")
	file = flag.String("file", "", "/path/to/filename")
	flag.Parse()

	if len(os.Args[1:]) <= 3 {
		fmt.Println("Usage: propgen [flags] \nFlags:")
		flag.PrintDefaults()
		fmt.Println("Examples:")
		fmt.Println("\tpropgen -label NIFI -render nifiproperties -file /opt/nifi/conf/nifi.properties")
		fmt.Println("\tpropgen -label HADOOP -render coresite -file /opt/hadoop/etc/hadoop/conf/core-site.xml")
		fmt.Println("\tpropgen -label HADOOP -render hdfssite -file /opt/hadoop/etc/hadoop/conf/hdfs-site.xml")
		fmt.Println("\tpropgen -label ZOOKEEPER -render zoocfg -file /opt/zookeeper/conf/zoo.cfg")
		fmt.Println("\tpropgen -label KUDU -render mastergflagfile -file /opt/kudu/conf/master.gflagfile")
		os.Exit(1)
	}
}

func main() {
	cmd.GenerateFile(*label, *render, *file)
}
