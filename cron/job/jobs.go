package job

import "log"

type Job func()

func PrintHelloWorld() {
	log.Println("Hello World!")
}
