package main

import (
	"fmt"
	"strconv"
	"strings"

	helper "github.com/wesleynepo/advent-of-code-2023/internal"
)

func partOne() {
    input, err := helper.ReadLines("input")

    if err != nil {
        panic(err)
    }

    response := [][2]string{}


    for _, line := range input {
        left := 0
        right := len(line) - 1

        output := [2]string{}


        for left <= right {
            if line[left] > 47 && line[left] < 58 {
                output[0] = fmt.Sprint(line[left])
            } else {
                left++
            }

            if line[right] > 47 && line[right] < 58 {
                output[1] = fmt.Sprint(line[right])
            } else {
                right--
            }

            if output[0] != "" && output[1] != "" {
                break
            }
        }
        response = append(response, output)

    }

    convertPrint(response)
}

func startsWith(line string) string {
    return check(line, strings.HasPrefix)
}

func endsWith(line string) string {
    return check(line, strings.HasSuffix)
}

func check(line string, fn func(string, string) bool) string {
    digits := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five":"5", "six":"6" , "seven":"7", "eight":"8", "nine": "9"}

    for key, value := range digits {
        if fn(line, key) {
            return value 
        }
    }

    return ""
}

func partTwo() {
    input, err := helper.ReadLines("input")

    if err != nil {
        panic(err)
    }

    response := [][2]string{}


    for _, line := range input {
        left := 0
        right := len(line) - 1

        output := [2]string{}


        for left <= right {
            if line[left] > 47 && line[left] < 58 {
                output[0] = string(line[left])
            } else {
                valid := startsWith(line[left:]) 

                if valid != "" {
                    output[0] = valid
                } else {
                    left++
                }
            }

            if line[right] > 47 && line[right] < 58 {
                output[1] = string(line[right])
            } else {
                valid := endsWith(line[:right+1]) 

                if valid != "" {
                    output[1] = valid
                } else {
                right--
                }
            }

            if output[0] != "" && output[1] != "" {
                break
            }
        }
        println(output[0], output[1])
        response = append(response, output)

    }
    convertPrint(response)

}

func convertPrint(response [][2]string) {
    sum := 0

    for _, pair := range response {
        value := pair[0] + pair[1]
        println(value)
        val, _ := strconv.Atoi(value) 
        sum += val
    }

    fmt.Println(sum)
}

func main() {
    partTwo()
}

