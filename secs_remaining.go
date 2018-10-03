package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	ui "uilive"
)

func errCheck(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func yearsToSecs(y int) int {
	return y * 3155695200
}

const LIFE_EXPECTANCY int = 259398145440

func main() {
	//get age
	fmt.Println("what is your age?")
	reader := bufio.NewReader(os.Stdin)
	text, e := reader.ReadString('\n')
	errCheck(e)
	text = strings.Replace(text, "\n", "", -1)
	age, e := strconv.Atoi(text)
	errCheck(e)

	//convert age to seconds
	age = yearsToSecs(age)

	//print remaining seconds
	writer := ui.New()
	writer.Start()
	for {
		//setup timer
		loop_timer := time.NewTimer(time.Second)

		//get remaining seconds
		remaining_secs := LIFE_EXPECTANCY - age

		//print
		fmt.Fprintf(writer, "           %d \n", remaining_secs)
		time.Sleep(time.Millisecond * 40)
		fmt.Fprintf(writer, "Remaining: %d \n", remaining_secs)

		//increment age
		age++

		//wait for timer
		<-loop_timer.C
	}
}
