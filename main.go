package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: incorrect number of arguments")
		return
	}

	s := ""
	for i := 0; i < len(os.Args[1]); i++ {
		if os.Args[1][i] == 10 {
			s += "\\" + "n"
		} else {
			s += string(os.Args[1][i])
		}
	}
	s += " "

	for i := 0; i < len(os.Args[1]); i++ {
		if (os.Args[1][i] > 126 || os.Args[1][i] < 32) && os.Args[1][i] != 10 {
			fmt.Println("Error: incorrect value input")
			return
		}
	}

	data, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: on reading file")
		return
	}

	if len(os.Args[1]) <= 0 {
		return
	}

	hash := MD5(string(data))

	if hash != "a51f800619146db0c42d26db3114c99f" {
		fmt.Println("Error: hash of file standard.txt has been changed")
		return
	}

	Replaced := strings.ReplaceAll(string(data), "\n\n", "\n")
	Replaced = strings.ReplaceAll(Replaced, "\\n", "\n")
	Splited := strings.Split(Replaced, "\n")

	ascii := make(map[byte]int)
	var q byte
	for q = 32; q <= 126; q++ {
		ascii[q] = (int(q)-32)*8 + 1
	}

	var newLine []int
	counter := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == 92 && s[i+1] == 'n' {
			newLine = append(newLine, i)
			counter++
			i++
		}
	}
	var result []string

	if counter == 0 {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(s)-1; j++ {
				result = append(result, Splited[ascii[s[j]]+i])
			}
			result = append(result, "\n")
		}
	} else {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(s)-1; j++ {
				if j+1 < (len(s)-1) && s[j] == 92 && s[j+1] == 'n' {
					break
				}
				result = append(result, Splited[ascii[s[j]]+i])
			}
			result = append(result, "\n")
		}
		for k := 0; k < len(newLine); k++ {
			for i := 0; i < 8; i++ {
				for j := newLine[k] + 2; j < len(s)-1; j++ {
					if j+1 < (len(s)-1) && s[j] == 92 && s[j+1] == 'n' {
						break
					}
					result = append(result, Splited[ascii[s[j]]+i])
				}
				result = append(result, "\n")
			}
		}
	}
	res := ""
	for i := 0; i < len(result); i++ {
		res += result[i]
	}

	Text := strings.ReplaceAll(res, "\n\n\n\n\n\n\n\n", "\n")

	counter_new := 0
	for i := 0; i < len(s)-1; i++ {
		if i+1 < (len(s)-1) && s[i] == 92 && s[i+1] == 'n' {
			counter_new++
			i++
		} else {
			counter_new = 0
			break
		}
	}

	if counter_new == 0 {
		fmt.Print(Text)
	} else if counter_new == (len(s)-1)/2 {
		for i := 0; i < counter_new; i++ {
			fmt.Println()
		}
	} else {
		fmt.Print(Text)
	}
}

func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}
