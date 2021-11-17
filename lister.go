// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Very simple tcp bandwidth test
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var version = ""

func listener() {
	fmt.Println("Bandwidth trasher listener - written by paul schou (github.com/pschou/bandwidth-trasher), version", version)
	fmt.Println("Loading environment settings")
	CONN_TYPE := getEnv("CONN_TYPE", "tcp")
	CONN_PORT := getEnv("CONN_PORT", "3333")
	CONN_HOST := getEnv("CONN_HOST", "")
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 10240)
	last_report := time.Now()
	var current_time time.Time
	var time_diff time.Duration
	var total, last_total uint64
	// Read the incoming connection into the buffer.
	for {
		c, err := conn.Read(buf)
		if total == 0 && string(buf[:4]) == "PULL" {
			fmt.Println("  Pull connection initiated")
			handleSend(conn)
			return
		}
		total += uint64(c)
		//fmt.Println("total", total, time_diff)
		current_time = time.Now()
		time_diff = current_time.Sub(last_report)
		if time_diff > 1e9 {
			fmt.Printf("  %+v %d %0.2fMbps (application layer speed)\n", conn.RemoteAddr(), total, float64(total-last_total)/float64(time_diff)*953.67)
			last_report = current_time
			last_total = total
		}
		if err != nil {
			fmt.Println("  Connection:", err.Error())
			break
		}
	}
	// Close the connection when you're done with it.
	conn.Close()
}
