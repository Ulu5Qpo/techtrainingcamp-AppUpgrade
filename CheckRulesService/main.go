package main

import "CheckRulesService/routers"

func main() {
	r := routers.SetupRouter()
	r.Run()
}
