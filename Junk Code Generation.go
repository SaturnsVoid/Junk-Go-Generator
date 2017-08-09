// Junk Code Generation project main.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	JunkCode string = ""

	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	vartypes    = [...]string{"byte", "bool", "int", "string"}
	retruntypes = [...]string{"bool", "int", "string"}
	booltypes   = [...]string{"true", "false"}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Junk Code Generation")

	JunkCode += `//Junk Code Generation` + "\r\n"
	JunkCode += `//https://github.com/SaturnsVoid` + "\r\n"

	for i := 0; i < randInt(5, 100); i++ {
		JunkCode += variableGeneration()
	}
	for i := 0; i < randInt(5, 100); i++ {
		JunkCode += functionGeneration()
	}

	createFileAndWriteData("JunkGoCode.txt", []byte(JunkCode))
	fmt.Println("Junk Code saved to `JunkGoCode.txt`")
}

func functionGeneration() string {
	var tmpint int = randInt(0, 11)
	var generated string = ""
	if tmpint > 5 {
		var value1 = "`" + RandStringRunes(randInt(3, 25)) + "`"
		var value2 = "`" + RandStringRunes(randInt(3, 25)) + "`"
		var tmpvar1 string = RandStringRunes(randInt(3, 25))
		var tmpvar2 string = RandStringRunes(randInt(3, 25))
		generated += "func " + RandStringRunes(randInt(3, 25)) + "() " + "{ " + "\r\n"
		generated += "var " + tmpvar1 + " string = " + value1 + "\r\n"
		generated += "var " + tmpvar2 + " string = " + value2 + "\r\n"
		generated += "if " + tmpvar1 + " == " + tmpvar2 + " {" + "\r\n"
		generated += "} else {" + "\r\n"
		generated += "}" + "\r\n"
		generated += "}" + "\r\n"
		return generated
	}
	var value1 = "`" + RandStringRunes(randInt(3, 25)) + "`"
	var value2 = "`" + RandStringRunes(randInt(3, 25)) + "`"
	var value3 = RandStringRunes(randInt(3, 25))
	var tmpvar1 string = RandStringRunes(randInt(3, 25))
	var tmpvar2 string = RandStringRunes(randInt(3, 25))
	var tmpvar3 string = ""
	var typ string = retruntypes[rand.Intn(len(retruntypes))]
	if typ == "string" {
		tmpvar3 = RandStringRunes(randInt(3, 25))
	} else if typ == "int" {
		tmpvar3 = strconv.Itoa(randInt(500, 999999))
	} else if typ == "bool" {
		tmpvar3 = booltypes[rand.Intn(len(booltypes))]
	}
	generated += "func " + RandStringRunes(randInt(3, 25)) + "() " + typ + "{ " + "\r\n"
	generated += "var " + tmpvar1 + " string = " + value1 + "\r\n"
	generated += "var " + tmpvar2 + " string = " + value2 + "\r\n"
	if typ == "string" {
		generated += "var " + value3 + " " + typ + " = `" + tmpvar3 + "` " + "\r\n"
	} else {
		generated += "var " + value3 + " " + typ + " = " + tmpvar3 + "\r\n"
	}
	generated += "if " + tmpvar1 + " == " + tmpvar2 + " {" + "\r\n"
	if typ == "int" {
		generated += value3 + " += " + strconv.Itoa(randInt(500, 999999)) + "\r\n"
	} else if typ == "bool" {
		generated += value3 + " = " + booltypes[rand.Intn(len(booltypes))] + "\r\n"
	} else if typ == "string" {
		generated += value3 + " += `" + RandStringRunes(randInt(3, 25)) + "`" + "\r\n"
	}
	generated += "} else {" + "\r\n"
	if typ == "int" {
		generated += value3 + " += " + strconv.Itoa(randInt(500, 999999)) + "\r\n"
	} else if typ == "bool" {
		generated += value3 + " = " + booltypes[rand.Intn(len(booltypes))] + "\r\n"
	} else if typ == "string" {
		generated += value3 + " += `" + RandStringRunes(randInt(3, 25)) + "`" + "\r\n"
	}
	generated += "}" + "\r\n"
	generated += "return " + value3 + "\r\n"
	generated += "}" + "\r\n"
	return generated
}

func variableGeneration() string {
	var tmptype string = vartypes[rand.Intn(len(vartypes))]
	var value string = ""
	if tmptype == "byte" {
		tmptype = ""
		value = "[]byte(" + "`" + RandStringRunes(randInt(3, 25)) + "`" + ")"
	} else if tmptype == "bool" {
		value = booltypes[rand.Intn(len(booltypes))]
	} else if tmptype == "int" {
		value = strconv.Itoa(randInt(500, 999999))
	} else if tmptype == "string" {
		value = "`" + RandStringRunes(randInt(3, 25)) + "`"
	}
	return "var " + RandStringRunes(randInt(3, 25)) + " " + tmptype + " = " + value + "\r\n"
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func createFileAndWriteData(fileName string, writeData []byte) error {
	fileHandle, err := os.Create(fileName)

	if err != nil {
		return err
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	writer.Write(writeData)
	writer.Flush()
	return nil
}
