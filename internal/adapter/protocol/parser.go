package protocol

import (
	"fmt"
	"redis-like-golang/internal/domain/command"
	"strings"
)

type Comand struct {
	Type command.Type
	Args []string
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}
func (p *Parser) ParseCommand(line string) (*Comand, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, fmt.Errorf("empty command")
	}

	parts := strings.Fields(line)
	if len(parts) == 0 {
		return nil, fmt.Errorf("invalid command format")
	}

	cmdType := command.Type(parts[0])
	if !command.IsValidCommandType(cmdType) {
		return nil, fmt.Errorf("invalid command type: %s", parts[0])
	}

	cmd := &Comand{
		Type: cmdType,
		Args: []string{},
	}

	if len(parts) > 1 {
		cmd.Args = parts[1:]
	}

	return cmd, nil
}

func (p *Parser) FormatResponse(response any) string {
	switch v := response.(type) {
	case string:
		return v
	case int, int64:
		return fmt.Sprintf(":%d\r\n", v)
	case bool:
		if v {
			return p.FormatOK()
		}
		return "FAIL"
	case error:
		return p.FormatError(v.Error())
	default:
		return fmt.Sprintf("%v\r\n", v)
	}
}

func (p *Parser) FormatOK() string {
	return "OK\r\n"
}

func (p *Parser) FormatError(msg string) string {
	return fmt.Sprintf("-ERR %s\r\n", msg)
}

func (p *Parser) FormatNill() string {
	return "nil\r\n"
}
