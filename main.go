package main

import serv "go_grpc/server"
func main() {
	g := serv.New()
	g.Start()
	g.WaitStop()
}
