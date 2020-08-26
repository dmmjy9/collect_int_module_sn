package main

import (
	"collect_int_module_sn/ssh"
)

func main() {
	//ssh.OpenScpServ()

	ssh.UploadFile("S6800-CMW710-SYSTEM-R2612P02H27.bin")

	//var wg sync.WaitGroup
	//ssh.Patch("S6800-CMW710-SYSTEM-R2612P02H27.bin", &wg)
	//wg.Wait()

	//ssh.Check()
}
