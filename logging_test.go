package main_test

import (
	"errors"
	"testing"

	logging "github.com/bkpeh/wgo"
)

func TestPrintErr(t *testing.T) {
	logging.PrintErr("Error Test", nil)

	e := errors.New("Error Test")

	logging.PrintErr("Error Test", e)
}

func TestPrintInfo(t *testing.T) {
	logging.PrintInfo("Info Test", nil)
}

func TestPrintDebug(t *testing.T) {
	logging.PrintDebug("Debug Test", nil)
}
