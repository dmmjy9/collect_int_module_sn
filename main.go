package main

import (
	"collect_int_module_sn/api/route"
)

func main() {
	r := route.InitRoute()
	r.Run(":8080")
}
