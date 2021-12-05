package lab2

import (
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	postfixExpression := ""
	partial := make([]byte, 8)
	for {
		count, err := (ch.Input).Read(partial)
		if err != nil{
			break
		}
		postfixExpression += string(partial[:count])
	}
	res, err := ConvertPostfixToPrefix(postfixExpression)
	if err != nil {
		return err;
	}
	ch.Output.Write([]byte(res))
	return nil
}