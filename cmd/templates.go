package cmd

const (
	ZooCfg           string = "zoocfg"
	CoreSite                = "coresite"
	HdfsSite                = "hdfssite"
	NifiProperties          = "nifiproperties"
	KuduMaster              = "kudumaster"
	KuduTserver             = "kudutserver"
	ImpalaDaemon            = "impaladaemon"
	ImpalaCatalog           = "impalacatalog"
	ImpalaStateStore        = "impalastatestore"
	ImpalaAdmission         = "impalaadmission"
	HiveSite				= "hivesite"
	MetastoreSite			= "metastoresite"
)

const coresite = "<?xml version=\"1.0\"?>\n" +
	"<?xml-stylesheet type=\"text/xsl\" href=\"configuration.xsl\"?>\n" +
	"<configuration>\n" +
	"{{range $key, $val := .}}\n" +
	"  <property>\n" +
	"      <name>{{$key}}</name>\n" +
	"      <value>{{$val}}</value>\n" +
	"  </property>\n" +
	"{{end}}\n" +
	"</configuration>\n"
const hdfssite = coresite

const zoocfg = "{{range $key, $val := .}}" +
	"{{$key}}={{$val}}\n" +
	"{{end}}"
const nifiproperties = zoocfg

const kudumaster = "{{range $key, $val := .}}" +
	"--{{$key}}={{$val}}\n" +
	"{{end}}"
const kudutserver = kudumaster
const impaladaemon = kudumaster
const impalacatalog = kudumaster
const impalastatestore = kudumaster
const impalaadmission = kudumaster
const hivesite = coresite
const metastoresite = coresite
