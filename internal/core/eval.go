package core

import (
	"errors"
	"fmt"
	"io"
)

func cmdPING(args []string) []byte {
	var buf []byte

	if len(args) > 1 {
		return Encode(errors.New("ERR wrong number of arguments for 'PING' command"), false)
	}

	if len(args) == 0 {
		buf = Encode("PONG", true)
	} else {
		buf = Encode(args[0], false)
	}

	return buf
}

func EvalAndResponse(cmd *Cmd, c io.ReadWriter) error {
	var res []byte

	switch cmd.Name {
	case "PING":
		res = cmdPING(cmd.Args)
	case "SET":
		res = cmdSET(cmd.Args)
	case "GET":
		res = cmdGET(cmd.Args)
	default:
		return errors.New(fmt.Sprintf("command not found: %s", cmd.Name))
	}

	_, err := c.Write(res)

	return err
}
