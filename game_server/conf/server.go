package conf


var Server struct{
	LogLevel    string //日志等级
	LogPath     string //日志路径
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string //tcp地址
	MaxConnNum  int    //最大连接数量
	ConsolePort int    //命令行端口
	ProfilePath string

	RedisAddrs struct { //redis
		LogAddr string
		DBAddr  string
	}

	PlatformID      int
	HttpServer      string //http服务器地址
	Mysql           string //mysql
}
