package cmd

const (
	ZooCfg         string = "zoocfg"
	CoreSite              = "coresite"
	HdfsSite              = "hdfssite"
	NifiProperties        = "nifiproperties"
)

var coresite = "<?xml version=\"1.0\"?>\n" +
	"<?xml-stylesheet type=\"text/xsl\" href=\"configuration.xsl\"?>\n" +
	"<configuration>\n" +
	"{{range $key, $val := .}}\n" +
	"  <property>\n" +
	"      <name>{{$key}}</name>\n" +
	"      <value>{{$val}}</value>\n" +
	"  </property>\n" +
	"{{end}}\n" +
	"</configuration>\n"

var zoocfg = "{{range $key, $val := .}}" +
	"{{$key}}={{$val}}\n" +
	"{{end}}"
