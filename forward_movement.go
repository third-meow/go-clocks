package main

import (
	"fmt"
	"time"
	ui "uilive"
)

func errCheck(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	writer := ui.New()
	writer.Start()
	for {
		//setup timer
		loop_timer := time.NewTimer(time.Second)

		now := time.Now()

		//print
		fmt.Fprintf(writer, "           %d \n", now.Unix())
		time.Sleep(time.Millisecond * 40)
		fmt.Fprintf(writer, "Time :     %d \n", now.Unix())

		//wait for timer
		<-loop_timer.C
	}
}
