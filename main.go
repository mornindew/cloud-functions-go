// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"math/rand"

	"./nodego"
)

func main() {
	flag.Parse()
	http.HandleFunc(nodego.HTTPTrigger, func(w http.ResponseWriter, r *http.Request) {
		processRequest(w, r)
//		fmt.Fprintln(w, "Hello, I'm native Go and generated the number", randomGenerator())
	})
	nodego.TakeOver()
}

func randomGenerator() int {
	return rand.Intn(1000)
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Everything Looks OK"))
}
