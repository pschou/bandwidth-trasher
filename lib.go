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
	"golang.org/x/crypto/salsa20/salsa"
	"net"
	"os"
	"time"
	"unsafe"
)

// Handles incoming requests.
func handleSend(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf, out := make([]byte, 10240), make([]byte, 10240)
	last_report := time.Now()
	var current_time time.Time
	var time_diff time.Duration
	var total, last_total uint64

	fmt.Println("Using Salsa20 stream cipher to generate randomness")
	counter := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var pcounter *uint64
	pcounter = (*uint64)(unsafe.Pointer(&counter[8]))
	var key = [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// Read the incoming connection into the buffer.
	for {
		//fmt.Printf("%+v\n", pcounter)
		*pcounter = *pcounter + 1
		//fmt.Println("calling salsa", len(buf))
		salsa.XORKeyStream(buf, out, &counter, &key)
		//fmt.Println("writing", len(buf))
		//fmt.Printf("writing %+v\n", counter, pcounter)
		c, err := conn.Write(buf)
		total += uint64(c)
		current_time = time.Now()
		time_diff = current_time.Sub(last_report)
		if time_diff > 1e9 {
			fmt.Printf("  %s %d %0.2fMbps (application layer speed)\n", conn.RemoteAddr(), total, float64(total-last_total)/float64(time_diff)*953.67)
			last_report = current_time
			last_total = total
		}
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
	}
	// Close the connection when you're done with it.
	conn.Close()
}

func getEnv(variable, def string) string {
	if os.Getenv(variable) != "" {
		def = os.Getenv(variable)
		fmt.Printf("  %s = %q\n", variable, def)
	} else {
		fmt.Printf("  %s = %q (default)\n", variable, def)
	}
	return def
}
