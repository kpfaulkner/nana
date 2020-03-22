package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/reiver/go-telnet"
)

type NanaCaller struct{}

// CallTELNET does all the good stuff :)
// First thing, calls to the "login" function. This sends the username and password to GrandMA2.
// If the login fails, the program will die completely.
func (c NanaCaller) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	login(w)
	time.Sleep(time.Duration(500) * time.Millisecond)

	fixtureFull(w, 1002, 1003)

	fixtureFull(w, 1002, 1003)
	fixtureFull(w, 1002, 1003)
	fixtureFull(w, 1002, 1003)
	fixtureFull(w, 1002, 1003)
	fixtureFull(w, 1002, 1003)
	fixtureFull(w, 1002, 1003)

	faderBrightness(w, 2, 0)
	time.Sleep(time.Duration(1000) * time.Millisecond)
	faderBrightness(w, 2, 100)

	/*
		fixtureFull(w, 1002, 1003)
		fixtureFull(w, 1002, 1003)
		groupFull(w, 38)

		// set channel 6 to full.
		channelFull(w, 6)

		faderBrightness(w, 2, 100)

		for i := 0; i < 10; i++ {
			faderBrightness(w, 3, i)
			faderBrightness(w, 4, i)
			time.Sleep(time.Duration(100) * time.Millisecond)
		}
	*/
}

// channelFull turns a group of channels to full brightness. No idea what that means, ask Will :)
func channelFull(writer telnet.Writer, channelNo int) error {
	template := "channel %d full\r\n"
	cmd := fmt.Sprintf(template, channelNo)

	// cmd should be "channel 6 full\r\n"
	return executeCommand(writer, cmd)
}

// fixtureFull turns a group of fixtures to full. No idea what that means, ask Will :)               fixture 1002 thru 1003 full
func fixtureFull(writer telnet.Writer, firstfixture int, secondfixture int) error {
	template := "fixture %d thru %d full\r\n"
	cmd := fmt.Sprintf(template, firstfixture, secondfixture)

	fmt.Printf("fixture command is %s\n", cmd)
	// cmd should be "channel 6 full\r\n"
	return executeCommand(writer, cmd)
}

// faderBrightness changes a particular fader to a particular brightness setting.
func faderBrightness(writer telnet.Writer, faderNo int, brightnessPercentage int) error {
	template := "fader %d At %d\r\n"
	cmd := fmt.Sprintf(template, faderNo, brightnessPercentage)

	// cmd should be "fader 3 At 55\r\n"
	return executeCommand(writer, cmd)
}

// group 38 full
func groupFull(writer telnet.Writer, group int) error {
	template := "group %d full\r\n"
	cmd := fmt.Sprintf(template, group)

	fmt.Printf("groupfull command in %s\n", cmd)
	// cmd should be "group 38 full\r\n"
	return executeCommand(writer, cmd)
}

// login will send username and password to grandMA
// if it fails, it will give some error messages and STOP the program completely
// (this is what log.Fatal below does)
func login(writer telnet.Writer) error {
	l := "Login \"telnet\" \"0872\"\r\n\r\n\r\n"
	length, err := writer.Write([]byte(l))
	if err != nil {
		fmt.Printf("login error %s\n", err.Error())
		log.Fatal("LOGIN ERROR %s\n", err.Error())
	}

	fmt.Printf("write %d chars\n", length)

	return nil
}

// executeCommand does the actual work. Sends the command string to GrandMA
func executeCommand(writer telnet.Writer, command string) error {
	_, err := writer.Write([]byte(command))
	if err != nil {
		fmt.Printf("executeCommand error %s\n", err.Error())
		return err
	}

	time.Sleep(200 * time.Millisecond)
	return nil
}

func main() {

	// ignore the next 3 lines.
	grandmaAddr := flag.String("server", "10.0.0.91", "IP/Host of Grand-MA server")
	flag.Parse()
	var caller telnet.Caller = NanaCaller{}

	// this says to "call" (ie connect) to the grandma server (10.0.0.91) port 30000
	// once it's connected, it will automatically run the CallTELNET function (top-ish of file).
	// We dont specifically TELL it to call CallTELNET, its just the "DialToAndCall" function below knows
	// to do this automatically.
	telnet.DialToAndCall((*grandmaAddr)+":30000", caller)

	// just has a sleep afterwards.
	time.Sleep(time.Duration(2) * time.Second)

	return

}
