package main

import (
	"flag"
	"fmt"
	"log"
)

func usage() {
	fmt.Println("PropGen Version: ", Version)
	fmt.Println("Usage: propgen [flags] \nFlags:")
	flag.PrintDefaults()
	fmt.Println("Examples:")
	fmt.Println("\tpropgen -label NIFI -render nifiproperties -file /opt/nifi/conf/nifi.properties")
	fmt.Println("\tpropgen -label HADOOP -render coresite -file /opt/hadoop/etc/hadoop/conf/core-site.xml")
	fmt.Println("\tpropgen -label HADOOP -render hdfssite -file /opt/hadoop/etc/hadoop/conf/hdfs-site.xml")
	fmt.Println("\tpropgen -label ZOOKEEPER -render zoocfg -file /opt/zookeeper/conf/zoo.cfg")
	fmt.Println("\tpropgen -label KUDU -render kudumaster -file /opt/kudu/conf/master.gflagfile")
}

func checkParametersIntegrity() {
	log.Println("checking parameters integrity")
	switch *label {
	case "NIFI":
		if *render != "nifiproperties" {
			log.Fatalf("%s needs render parameter: nifiproperties", *label)
		}
	case "HADOOP":
		if *render != "coresite" && *render != "hdfssite" {
			log.Println(*render)
			log.Fatalf("%s needs the render parameter: coresite, hdfssite", *label)
		}
	case "ZOOKEEPER":
		if *render != "zoocfg" {
			log.Fatalf("%s needs render parameter: zoocfg", *label)
		}
	case "KUDU":
		if *render != "coresite" && *render != "hdfssite" {
			log.Println(*render)
			log.Fatalf("%s needs the render parameter: coresite, hdfssite", *label)
		}
	default:
		log.Fatal("Unspected Label: ", *label)
	}
}