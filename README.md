# propgen
Container Property File Generator for using in Kubernetes Deployments

```
$ ./propgen 
Usage: propgen [flags] 
Flags:
  -file string
        /path/to/filename
  -label string
        NIFI|HADOOP|ZOOKEEPER|KUDU
  -render string
        coresite|hdfssite|nifiproperties|zoocfg
Examples:
        propgen -label NIFI -render nifiproperties -file /opt/nifi/conf/nifi.properties
        propgen -label HADOOP -render coresite -file /opt/hadoop/etc/hadoop/conf/core-site.xml
        propgen -label HADOOP -render hdfssite -file /opt/hadoop/etc/hadoop/conf/hdfs-site.xml
        propgen -label ZOOKEEPER -render zoocfg -file /opt/zookeeper/conf/zoo.cfg
        propgen -label KUDU -render mastergflagfile -file /opt/kudu/conf/master.gflagfile
```
