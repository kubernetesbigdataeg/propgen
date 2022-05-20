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
Examples:
        propgen -label NIFI -file /opt/nifi/conf/nifi.properties
        propgen -label HADOOP -file /opt/hadoop/etc/hadoop/conf/core-site.xml
        propgen -label ZOOKEEPER -file /opt/zookeeper/conf/zoo.cfg
```
