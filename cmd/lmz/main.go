package main

import (
	"fmt"
	"mp/lmz/pkg/auth"
	"mp/lmz/pkg/config"
	"mp/lmz/pkg/lmz"
	"os"
)

type Op int

const (
	_ Op = iota
	Status
	On
	Off
)

func main() {
	op := MustGetOp()

	c := config.MustRead()
	t := must(auth.GetToken(c))

	l := lmz.New(c, t)

	output := "OK"

	switch op {
	case On:
		check(l.TurnOn())
	case Off:
		check(l.TurnOff())
	default:
		status := must(l.Status())
		localTime := status.Received.Local()
		output = "Status as of " + localTime.String() + ": " + status.MachineStatus
	}

	fmt.Println(output)
}

func MustGetOp() Op {
	showUsage := len(os.Args) > 2 || (len(os.Args) > 1 && os.Args[1] != "on" && os.Args[1] != "off")
	if showUsage {
		fmt.Println("Usage: ", os.Args[0], "[on|off]")
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "on":
			return On
		case "off":
			return Off
		}
	}

	return Status
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func must[T any](thing T, err error) T {
	check(err)
	return thing
}
