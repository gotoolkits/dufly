package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/GeertJohan/go.linenoise"
	"github.com/gotoolkits/dufly/common"
)

const (
	root = "/"
)

var rootWorkPath = "yunpan_dufs"
var currentWorkPath = "/"
var archTag = "$ "

func main() {
	fmt.Println("Welcome to running the dufly console to manage the yunpan files! \nCommand list:")
	writeHelp()
	for {
		str, err := linenoise.Line(rootWorkPath + currentWorkPath + archTag)
		if err != nil {
			if err == linenoise.KillSignalError {
				quit()
			}
			fmt.Printf("Unexpected error happen: %s\n", err)
			quit()
		}
		fields := strings.Fields(str)
		addToHistory(str)

		// check if there is any valid input at all
		if len(fields) == 0 {
			writeUnrecognized()
			continue
		}

		// switch on the command
		switch fields[0] {
		case "help":
			writeHelp()
		case "echo":
			fmt.Printf("echo: %s\n\n", str[5:])
		case "clear":
			linenoise.Clear()
		case "ls":
			listFiles(str)
		case "cd":
			changeDirctory(str)
		case "pwd":
			fmt.Println(getWorkPath())
		case "multiline":
			fmt.Println("Setting dufly to multiline")
			linenoise.SetMultiline(true)
		case "singleline":
			fmt.Println("Setting dufly to singleline")
			linenoise.SetMultiline(false)
		case "printKeyCodes":
			linenoise.PrintKeyCodes()
		case "save":
			if len(fields) != 2 {
				fmt.Println("Error. Expecting 'save <filename>'.")
				continue
			}
			err := linenoise.SaveHistory(fields[1])
			if err != nil {
				fmt.Printf("Error on save: %s\n", err)
			}
		case "load":
			if len(fields) != 2 {
				fmt.Println("Error. Expecting 'load <filename>'.")
				continue
			}
			err := linenoise.LoadHistory(fields[1])
			if err != nil {
				fmt.Printf("Error on load: %s\n", err)
			}
		case "quit":
			quit()
		case "exit":
			quit()
		default:
			writeUnrecognized()
		}
	}
}

func quit() {
	fmt.Println("Thanks for running the dufly console.")
	fmt.Println("")
	os.Exit(0)
}

func writeHelp() {
	fmt.Println("help                    write this message")
	fmt.Println("echo ...                echo the arguments")
	fmt.Println("clear                   clear the screen")
	fmt.Println("login                   login to YunPan get Auth Session")
	fmt.Println("ls [path]               list the file or dirctory")
	fmt.Println("cd [path]               change the work path")
	fmt.Println("pwd                     show the current work path")
	fmt.Println("upload [local][dst]     upload the localfile to YunPan")
	fmt.Println("delete [path]           delete the YunPan files")
	//	fmt.Println("multiline           set dufly to multiline")
	//	fmt.Println("singleline          set dufly to singleline")
	//	fmt.Println("save <filename>     save the command history to file")
	//	fmt.Println("load <filename>     load the command history from file")
	fmt.Println("quit/exit               stop the program")
	fmt.Println("")
}

func writeUnrecognized() {
	fmt.Println("Unkown command to reconize. please use 'help' info.")
}

func changeDirctory(cmd string) {

	fields := strings.Fields(cmd)

	if len(fields) < 2 {
		fmt.Println("No arguments set,current dirctory is ", getWorkPath())
		return
	}
	workPath := fields[1]

	if isBack(workPath) {
		return
	}

	if !isDirctoryExist(getWorkPath() + "/" + workPath) {
		fmt.Println("Dirctory No Found ,Please Check your input.")
		return
	}

	if isAbsolutePath(workPath) {
		fmt.Println("isAbsolutePath is true")
		currentWorkPath = workPath
		return
	}

	currentWorkPath = currentWorkPath + "/" + workPath
	fmt.Println(currentWorkPath)
}

func getWorkPath() string {
	return currentWorkPath
}

func isDirctoryExist(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func listFiles(path string) {
	fields := strings.Fields(path)
	if len(fields) == 1 {
		if currentWorkPath == "" {
			fmt.Println(common.CommonandLs("/"))
			return
		}
		fmt.Println(common.CommonandLs(currentWorkPath))
		return
	}
	if len(fields) != 2 {
		fmt.Println("Arguments is promblem??")
		return
	}
	fmt.Println(common.CommonandLs(currentWorkPath + fields[1]))
}

func addToHistory(cmd string) {
	err := linenoise.AddHistory(cmd)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func isBack(cmd string) bool {
	var dlist []string = nil
	if cmd == ".." {
		if isRootPath() {
			return true
		}
		dlist = strings.Split(currentWorkPath, "/")
		end := len(dlist) - 1

		currentWorkPath = strings.Join(dlist[:end], "/")

		isRootPath()
		return true
	}
	if cmd == "/" {
		currentWorkPath = "/"
		return true
	}
	return false
}

func isAbsolutePath(path string) bool {

	path = strings.TrimSpace(path)
	if string(path[0]) == "/" {
		return true
	}
	return false
}

func isRootPath() bool {
	if currentWorkPath == "" {
		fmt.Println("current path is already / ")
		currentWorkPath = "/"
		return true
	}

	return false
}
