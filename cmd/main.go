package main

import "txp/rabbitmq/app"

func main() {
	app.Send()
	app.Receive()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
