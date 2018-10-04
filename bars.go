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

func create_bar(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		s = s + "*"
	}
	return s
}

func main() {
	writer := ui.New()
	writer.Start()
	for {
		//setup timer
		loop_timer := time.NewTimer(time.Second)

		now := time.Now()
		hours_bar := create_bar(now.Hour())
		mins_bar := create_bar(now.Minute())
		secs_bar := create_bar(now.Second())

		//print
		fmt.Fprintf(writer, "Hours |%s\n", hours_bar)
		fmt.Fprintf(writer, "Mins  |%s\n", mins_bar)
		fmt.Fprintf(writer, "Secs  |%s\n", secs_bar)

		//wait for timer
		<-loop_timer.C
	}
}
