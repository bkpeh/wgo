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

//PrintLog for logging. s is either info,err,debug
func printlog(logtype string, logstr string, e error) {
	fopen, ferr := os.OpenFile(`\logs\logfile.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logging := log.New(fopen, s, log.LstdFlags)

	if ferr != nil {
		log.Println(logerr, e)
	}

	//Check if there is any error passed in
	if e != nil {
		log.Println(logtype, logstr, e)
	} else {
		log.Println(logtype, logstr)
	}

}

func PrintErr(s string, e error) {
	PrintLog(logerr, s, e)
}
