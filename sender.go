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
)

func sender() {
	fmt.Println("Bandwidth trasher sender - written by paul schou (github.com/pschou/bandwidth-trasher), version", version)
	fmt.Println("Loading environment settings")
	CONN_TYPE := getEnv("CONN_TYPE", "tcp")
	CONN_PORT := getEnv("CONN_PORT", "3333")
	CONN_HOST := getEnv("CONN_HOST", "localhost")
	// Listen for incoming connections.
	l, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Connecting to " + CONN_HOST + ":" + CONN_PORT)

	handleSend(l)
}
