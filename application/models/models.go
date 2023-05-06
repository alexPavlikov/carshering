package models

var Config Cfg

type Cfg struct {
	ServerHost  string
	ServerPort  string
	PgHost      string
	PgPort      string
	PgUser      string
	PgPass      string
	PgName      string
	Data        string
	Assets      string
	Html        string
	MainEmail   string
	MainOffice  string
	CompanyName string
	IndexTitle  string
	Slogan      string
	Year        string
}
