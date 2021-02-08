package main

import (
	"log"
	"os"
)

const (
	loginfo  = "[INFO]"
	logerr   = "[ERROR]"
	logdebug = "[DEBUG]"
)

func printlog(logtype string, logstr string, e error) {
	fopen, ferr := os.OpenFile(`logs\logfile.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if ferr != nil {
		log.Println(logerr, e)
	}

	defer fopen.Close()

	logging := log.New(fopen, logtype, log.LstdFlags)

	//Check if there is any error passed in
	if e != nil {
		logging.Println(logstr, e)
	} else {
		logging.Println(logstr)
	}

}

//PrintErr is to print with [ERROR] prefix
func PrintErr(s string, e error) {
	printlog(logerr, s, e)
}

//PrintInfo is to print with [INFO] prefix
func PrintInfo(s string, e error) {
	printlog(loginfo, s, e)
}

//PrintDebug is to print with [DEBUG] prefix
func PrintDebug(s string, e error) {
	printlog(logdebug, s, e)
}
