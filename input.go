package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

// 1つintを返す
// input 1
// output 1
func NextInt() int {
	ret, err := strconv.Atoi(NextString())
	if err != nil {
		panic(err)
	}
	return ret
}

// 1つのstringを返す
// input 1
// output "1"
func NextString() string {
	sc.Scan()
	ret := sc.Text()
	return ret
}

// Return a int slice line
/*
input
1
1
1
output
[1 1 1]
*/
func OneLineReadInt(n int) []int {
	l := []int{}
	for i := 0; i < n; i++ {
		l = append(l, NextInt())

	}
	return l

}

// Return a string slice line
/*
input
1
1
1
output
["1" "1" "1"]
*/
func OneLineReadString(n int) []string {
	l := []string{}
	for i := 0; i < n; i++ {
		l = append(l, NextString())
	}
	return l
}

// Return a string slice line
/*
input
1 1 1
output
["1" "1" "1"]
*/
func LineReadString() []string {
	l := strings.Split(NextString(), " ")
	return l

}

// Return a int slice line
/*
input
1 1 1
output
[1 1 1]
*/
func LineReadInt() []int {
	l := LineReadString()
	x := []int{}
	for _, v := range l {
		tmp, _ := strconv.Atoi(v)
		x = append(x, tmp)
	}
	return x
}

// Return string slices lines
/*
input
1 1 1
1 1 1
1 1 1
output
[
	["1" "1" "1"]
	["1" "1" "1"]
	["1" "1" "1"]
]
*/
func LinesReadString(n int) [][]string {
	l := [][]string{}
	for i := 0; i < n; i++ {
		l = append(l, LineReadString())
	}
	return l
}

// Return int slices lines
/*
input
1 1 1
1 1 1
1 1 1
output
[
	[1 1 1]
	[1 1 1]
	[1 1 1]
]
*/
func LinesReadInt(n int) [][]int {
	l := [][]int{}
	for i := 0; i < n; i++ {
		l = append(l, LineReadInt())
	}
	return l
}

// Return int slices lines
/*
input
1 2
3 4
5 6
output
["1" "3" "5"] ["2" "4" "6"]
*/
func TwoLinesReadString(n int) ([]string, []string) {
	xy := LinesReadString(n)
	x := []string{}
	y := []string{}
	for _, v := range xy {
		x = append(x, v[0])
		y = append(y, v[1])
	}
	return x, y
}

// Return int slices lines
/*
input
1 2
3 4
5 6
output
["1" "3" "5"] ["2" "4" "6"]
*/
func TwoLinesReadInt(n int) ([]int, []int) {
	xy := LinesReadInt(n)
	x := []int{}
	y := []int{}
	for _, v := range xy {
		x = append(x, v[0])
		y = append(y, v[1])
	}
	return x, y
}
