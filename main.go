package main

import appsetup "github.com/ghaneakbar4/ApiGolang/AppSetUp"

//نقطه شروع برنامه
func main()  {
	var obj appsetup.App
	obj.DatabaseConection()
	obj.Routes()
	obj.Run()
}