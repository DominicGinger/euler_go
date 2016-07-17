package main

import (
	"io/ioutil"
	"sort"
	"strings"
)

func alphabeticalize(names_arr []string) []string {
	sort.Strings(names_arr)
	return names_arr
}

func count(name string, idx int) int {
	sum := 0
	for _, v := range []byte(name) {
		sum += (int(v) - 64)
	}
	return sum * idx
}

func p022() int {
	file, _ := ioutil.ReadFile("p022_names.txt")
	names := strings.Trim(string(file), "\"")
	names_arr := strings.Split(names, "\",\"")
	total := 0
	for i, v := range alphabeticalize(names_arr) {
		total += count(v, i+1)
	}
	return total
}
