package main

import (
	"flag"
	"github.com/kpfaulkner/nana/pkg"
	"time"

	"github.com/reiver/go-telnet"

)


func main() {

	// ignore the next 3 lines.
	grandmaAddr := flag.String("server", "10.0.0.91", "IP/Host of Grand-MA server")
	flag.Parse()
	var caller telnet.Caller = pkg.NanaCaller{}

	// this says to "call" (ie connect) to the grandma server (10.0.0.91) port 30000
	// once it's connected, it will automatically run the CallTELNET function (top-ish of file).
	// We dont specifically TELL it to call CallTELNET, its just the "DialToAndCall" function below knows
	// to do this automatically.
	telnet.DialToAndCall((*grandmaAddr)+":30000", caller)

	// just has a sleep afterwards.
	time.Sleep(time.Duration(2) * time.Second)

	return

}
