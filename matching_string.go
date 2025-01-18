package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMatchingStrings() string {
	var n int
	fmt.Print("Masukkan jumlah string: ")
	fmt.Scanln(&n)

	stringsList := make(map[string][]int)
	reader := bufio.NewReader(os.Stdin)

	var identicFound string

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan string ke-%d: ", i+1)
		input, _ := reader.ReadString('\n')
		stringsList[input] = append(stringsList[input], i+1)
		if len(stringsList[input]) > 1 && identicFound == "" {
			identicFound = input
		}
	}

	if value, ok := stringsList[identicFound]; ok {
		strList := make([]string, len(value))

		for i, v := range value {
			strList[i] = strconv.Itoa(v)
		}

		return strings.Join(strList, " ")
	}

	return "false"
}
