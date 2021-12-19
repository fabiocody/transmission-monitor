package main

type ArgsType struct {
	Debug    bool   `arg:"env" help:"activate debug logs"`
	Database string `arg:"env" default:".transmission-monitor.sqlite3" help:"database file"`
	Host     string `arg:"env" default:"localhost" help:"transmission host"`
	Port     uint16 `arg:"env" default:"9091" help:"transmission port"`
	Username string `arg:"env" help:"transmission username"`
	Password string `arg:"env" help:"transmission password"`
	Https    bool   `arg:"env" default:"false" help:"enable https when connecting to transmission"`
	ApiToken string `arg:"required,env:API_TOKEN" help:"telegram api token"`
	ChatId   int64  `arg:"required,env:CHAT_ID" help:"telegram chat id"`
	Daemon   bool   `arg:"env" help:"Daemonize (the check will run every minute)"`
}

var Args ArgsType
