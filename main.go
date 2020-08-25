package main

import (
	_ "jizhangji/boot"
	_ "jizhangji/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
