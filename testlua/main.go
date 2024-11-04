package main

import (
	"log"

	lua_debugger "github.com/gavin-wuz/gopherlua-debugger"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	lua_debugger.Preload(L)

	err := L.DoFile("main.lua")
	log.Println(err)
}
