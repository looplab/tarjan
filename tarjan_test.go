// Copyright (c) 2013 - Max Persson <max@looplab.se>
// Copyright (c) 2010-2013 - Gustavo Niemeyer <gustavo@niemeyer.net>
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

package tarjan

import (
	"fmt"
	"testing"
)

func TestTypeString(t *testing.T) {
	graph := make(map[interface{}][]interface{})
	graph["1"] = []interface{}{"2"}
	graph["2"] = []interface{}{"3"}
	graph["3"] = []interface{}{"1"}
	graph["4"] = []interface{}{"2", "3", "5"}
	graph["5"] = []interface{}{"4", "6"}
	graph["6"] = []interface{}{"3", "7"}
	graph["7"] = []interface{}{"6"}
	graph["8"] = []interface{}{"5", "7", "8"}

	output := Connections(graph)
	if output[0][0] != "3" {
		t.FailNow()
	}
	if output[0][1] != "2" {
		t.FailNow()
	}
	if output[0][2] != "1" {
		t.FailNow()
	}
	if output[1][0] != "7" {
		t.FailNow()
	}
	if output[1][1] != "6" {
		t.FailNow()
	}
	if output[2][0] != "5" {
		t.FailNow()
	}
	if output[2][1] != "4" {
		t.FailNow()
	}
	if output[3][0] != "8" {
		t.FailNow()
	}
}

func TestTypeInt(t *testing.T) {
	graph := make(map[interface{}][]interface{})
	graph[1] = []interface{}{2}
	graph[2] = []interface{}{3}
	graph[3] = []interface{}{1}
	graph[4] = []interface{}{2, 3, 5}
	graph[5] = []interface{}{4, 6}
	graph[6] = []interface{}{3, 7}
	graph[7] = []interface{}{6}
	graph[8] = []interface{}{5, 7, 8}

	output := Connections(graph)
	if output[0][0] != 3 {
		t.FailNow()
	}
	if output[0][1] != 2 {
		t.FailNow()
	}
	if output[0][2] != 1 {
		t.FailNow()
	}
	if output[1][0] != 7 {
		t.FailNow()
	}
	if output[1][1] != 6 {
		t.FailNow()
	}
	if output[2][0] != 5 {
		t.FailNow()
	}
	if output[2][1] != 4 {
		t.FailNow()
	}
	if output[3][0] != 8 {
		t.FailNow()
	}
}

func TestTypeMixed(t *testing.T) {
	graph := make(map[interface{}][]interface{})
	graph[1] = []interface{}{2}
	graph[2] = []interface{}{"3"}
	graph["3"] = []interface{}{1}
	graph[4] = []interface{}{2, "3", 5}
	graph[5] = []interface{}{4, "6"}
	graph["6"] = []interface{}{"3", 7}
	graph[7] = []interface{}{"6"}
	graph[8] = []interface{}{5, 7, 8}

	output := Connections(graph)
	if output[0][0] != "3" {
		t.FailNow()
	}
	if output[0][1] != 2 {
		t.FailNow()
	}
	if output[0][2] != 1 {
		t.FailNow()
	}
	if output[1][0] != 7 {
		t.FailNow()
	}
	if output[1][1] != "6" {
		t.FailNow()
	}
	if output[2][0] != 5 {
		t.FailNow()
	}
	if output[2][1] != 4 {
		t.FailNow()
	}
	if output[3][0] != 8 {
		t.FailNow()
	}
}

func TestEmptyGraph(t *testing.T) {
	graph := make(map[interface{}][]interface{})

	output := Connections(graph)
	if len(output) != 0 {
		t.FailNow()
	}
}

func TestSingleVertex(t *testing.T) {
	graph := make(map[interface{}][]interface{})
	graph["1"] = []interface{}{}

	output := Connections(graph)
	if output[0][0] != "1" {
		t.FailNow()
	}
}

func TestSingleVertexLoop(t *testing.T) {
	graph := make(map[interface{}][]interface{})
	graph["1"] = []interface{}{"1"}

	output := Connections(graph)
	if output[0][0] != "1" {
		t.FailNow()
	}
}

func TestMultipleVertexLoop(t *testing.T) {
	graph := make(map[interface{}][]interface{})
	graph["1"] = []interface{}{"2"}
	graph["2"] = []interface{}{"3"}
	graph["3"] = []interface{}{"1"}

	output := Connections(graph)
	if output[0][0] != "3" {
		t.FailNow()
	}
	if output[0][1] != "2" {
		t.FailNow()
	}
	if output[0][2] != "1" {
		t.FailNow()
	}
}

func ExampleConnections() {
	graph := make(map[interface{}][]interface{})
	graph["1"] = []interface{}{"2"}
	graph["2"] = []interface{}{"3"}
	graph["3"] = []interface{}{"1"}
	graph["4"] = []interface{}{"2", "3", "5"}
	graph["5"] = []interface{}{"4", "6"}
	graph["6"] = []interface{}{"3", "7"}
	graph["7"] = []interface{}{"6"}
	graph["8"] = []interface{}{"5", "7", "8"}

	output := Connections(graph)
	fmt.Println(output)

	// Output:
	// [[3 2 1] [7 6] [5 4] [8]]
}
