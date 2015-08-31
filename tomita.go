package tomita

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

// Parser это конфиг, в котором сохраняется путь до томиты и вашего конфига
type Parser struct {
	execPath string
	confPath string
	debug    bool
}

// New создает инстанс парсера
func New(execPath, confPath string) (Parser, error) {
	tp := Parser{
		execPath,
		confPath,
		false,
	}
	return tp, nil
}

// Debug устанавливает многословный режим
func (tp *Parser) Debug(debug bool) {
	tp.debug = debug
}

// SetDebug is deprecated
func (tp *Parser) SetDebug(debug bool) {
	log.Printf("SetDebug is deprecated. Use tp.Debug(%t).", debug)
}

// Run запускает парсер и считывает вывод
func (tp *Parser) Run(text string) (string, error) {
	tp.debugMsg("Run Tomita")

	cmd := exec.Command(tp.execPath, tp.confPath)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Panic(err)
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		log.Panic(err)
		return "", err
	}

	tp.debugMsg("Write input")
	_, err = stdin.Write([]byte(text))
	if err != nil {
		panic("Tomita probably hasn't waited for the input. It can be caused by an error in config.proto. " +
			err.Error())
	}
	stdin.Close()

	tp.debugMsg("Read output")
	output, err := ioutil.ReadAll(stdout)
	panicOnErr(err)

	tp.debugMsg(fmt.Sprintf("Output: \"%s\"", string(output)))

	return string(output), nil
}

func (tp *Parser) debugMsg(message string) {
	if tp.debug {
		log.Print(message)
	}
}

func panicOnErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
