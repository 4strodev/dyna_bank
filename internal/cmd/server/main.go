package main

import "log"

func main() {
	log.Fatal(app.Listen(":3000"))
}
