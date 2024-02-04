package initialize

// Run int
func Run() {
	LoadConfig()
	Redis()
	Mysql()
	Router()
}
