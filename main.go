package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"reflect"
	"strconv"
	"math"
)
var i int = 0
var numbers_list = make([][]string,100)
//var numbers_list = make([]string,100)
var zero = [][]string{
	{" ","_"," "},
	{"|"," ","|"},
	{"|","_","|"},
}
var one = [][]string{
	{" "," "," "},
	{" "," ","|"},
	{" "," ","|"},

}
var two = [][]string{
	{" ","_"," "},
	{" ","_","|"},
	{"|","_"," "},
}

var tree = [][]string{
	{" ","_"," "},
	{" ","_","|"},
	{" ","_","|"},
}
var foor = [][]string{
	{" "," "," "},
	{"|","_","|"},
	{" "," ","|"},
}
var five = [][]string{
	{" ","_"," "},
	{"|","_"," "},
	{" ","_","|"},
}
var six = [][]string{
	{" ","_"," "},
	{"|","_"," "},
	{"|","_","|"},
}
var seven = [][]string{
	{" ","_"," "},
	{" "," ","|"},
	{" "," ","|"},
}
var eight = [][]string{
	{" ","_"," "},
	{"|","_","|"},
	{"|","_","|"},
}
var nine = [][]string{
	{" ","_"," "},
	{"|","_","|"},
	{" ","_","|"},
}
func printOCR(orc [][]string){
	for i, _ := range orc {
		for j,_ := range orc[i] {
			fmt.Print(orc[i][j])
		}
		fmt.Println("")
	}
}

func compare(number [][]string) string{
	if reflect.DeepEqual(number,zero) {
		return "0"
	}
	if reflect.DeepEqual(number,one) {
		return "1"
	}
	if reflect.DeepEqual(number,two) {
		return "2"
	}
	if reflect.DeepEqual(number,tree) {
		return "3"
	}
	if reflect.DeepEqual(number,foor) {
		return "4"
	}
	if reflect.DeepEqual(number,five) {
		return "5"
	}
	if reflect.DeepEqual(number,six) {
		return "6"
	}
	if reflect.DeepEqual(number,seven) {
		return "7"
	}
	if reflect.DeepEqual(number,eight) {
		return "8"
	}
	if reflect.DeepEqual(number,nine) {
		return "9"
	}
	return"?"
}
func checksum(number string) string{
	sum := 0 
	n := strings.Split(number,"")
	for i:=0; i < len(n); i++ {
		val, _ := strconv.Atoi(n[i])
		sum += val*(len(n)-i)
	}
	if math.Mod(float64(sum), 11) == 0.{
		return ""
	}else{
		return "ERR"
	}
}
func convert(line [][]string){
	fmt.Println("begin converting ord code")
	fmt.Println("enter orc code is : ")
	printOCR(line)
	orc_converted := ""
	orc := make([][]string,3)
	for i:=0 ; i<27; i = i+3 {
		orc[0] = make([]string,1)
		orc[0] = strings.Split(line[0][0],"")[i:i+3]
		orc[1] = make([]string,1)
		orc[1] = strings.Split(line[1][0],"")[i:i+3]
		orc[2] = make([]string,1)
		orc[2] = strings.Split(line[2][0],"")[i:i+3]
		orc_converted +=compare(orc)
	}
	if strings.Contains("orc_converted", "?"){
		fmt.Println(orc_converted, " ", "ILL")
	}else {
		fmt.Println(orc_converted, " ", checksum(orc_converted))
	}
	
	

	
}

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanLines(data, atEOF)
    if err == nil && token != nil {
    	numbers_list[i] = make([]string,1)
        numbers_list[i][0] = string(token)
        i = i+1
    }	
    return
}

func main() {
	f, err := os.Open("bankOCR")
    if err != nil {
		panic(err)
	}
	
	scanner := bufio.NewScanner(f)
	scanner.Split(split)
	
	for scanner.Scan() {
		scanner.Text()
	}
	for i:=0; i<len(numbers_list); i=i+4 {
		convert(numbers_list[i:i+4])
	}
	
	
}
