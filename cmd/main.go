package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/reiver/go-telnet"
	"io"
	"net"
	"time"
)

// change

type NanaCaller struct{}

func readLoop( r telnet.Reader) error {
	bytes := make([]byte,1,1)

	for {
		n, err := r.Read(bytes)
		if err == io.EOF {
			// EOF.
			fmt.Printf("COMMS ENDED!\n")
			return err
		}

		if err != nil {
			// not EOF but real error!
			fmt.Printf("ERROR!!!! %s\n", err.Error())
			return err
		}

		// if read 0 bytes, sleep for a bit.
		if n == 0 {
			time.Sleep(time.Duration(500) * time.Millisecond)
		} else {
			fmt.Printf(string(bytes[:len(bytes)]))
		}
	}
}

func (c NanaCaller) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {

	go readLoop(r)

	for {
		time.Sleep(time.Duration(1) * time.Second)
	}

}

func readData(conn net.Conn) (string, error) {
	timeoutDuration := 1 * time.Second
	bufReader := bufio.NewReader(conn)

	totalString := ""
	for {
		// Set a deadline for reading. Read operation will fail if no data
		// is received after deadline.
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		// Read tokens delimited by newline
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			//fmt.Println(err)
			return totalString, nil
		}

		totalString += string(bytes)
		///fmt.Printf("%s", bytes)
	}
}

func readSplashScreen( conn net.Conn) error {

	ss,_ := readData(conn)

	fmt.Printf("splashscreen\n\n%s\n", ss)

	return nil
	timeoutDuration := 2 * time.Second
	bufReader := bufio.NewReader(conn)

	for {
		// Set a deadline for reading. Read operation will fail if no data
		// is received after deadline.
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		// Read tokens delimited by newline
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			//fmt.Println(err)
			return nil
		}

		fmt.Printf("%s", bytes)
	}

}

func doLogin(conn net.Conn ) error {
	// do the login!
	//loginString := `Login "telnet" 0872`

	//fmt.Fprintf(conn, loginString)
	fmt.Fprintf(conn, "\n")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: "+message)

	return nil

	/*
	fmt.Printf("login string: %s\n", loginString)
	// send to socket
	//fmt.Fprintf(conn, loginString+"\n")
	conn.Write([]byte( loginString))
	conn.Write([]byte( "\n"))
	// listen for reply

	time.Sleep(time.Duration(2)*time.Second)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)

	return nil
	time.Sleep(time.Duration(2)*time.Second)
	lr, _ := readData(conn)
	fmt.Printf("login reply %s\n", lr)
	//message, _ := bufio.NewReader(conn).ReadString('\n')
	//fmt.Printf("login reply %s\n", message)
*/
  return nil
}

func testidea() error {
	conn, _ := telnet.DialTo("rainmaker.wunderground.com:23")

	data := make([]byte,1,1)

	totalCount := 0
	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Printf("error on read %s\n", err.Error())
			return err
		}

		fmt.Printf("read %d bytes\n", n)
		totalCount += n
		fmt.Printf("read total %d bytes\n", totalCount)
		fmt.Printf("raw char is %v\n", data)
		fmt.Printf("char is %s\n", string(data))

	}
}

func main() {

	grandmaAddr := flag.String("server", "rainmaker.wunderground.com", "IP/Host of Grand-MA server")
	//username := flag.String("username", "", "Grand-MA username")
	//password := flag.String("password", "", "Grand-MA password")
	//verbose := flag.Bool("verbose", false, "spamageddon")

	flag.Parse()


	// connect to this socket
	//conn, _ := net.Dial("tcp", (*grandmaAddr)+":23")

	var caller telnet.Caller = NanaCaller{}

	//@TOOD: replace "example.net:5555" with address you want to connect to.
	telnet.DialToAndCall((*grandmaAddr)+":23", caller)


  //readSplashScreen(conn)
	//readSplashScreen(conn)
	time.Sleep(time.Duration(2)*time.Second)
	//doLogin(conn)

	return

}
