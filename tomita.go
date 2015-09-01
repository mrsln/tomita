// tomita это обертка для Томита-парсер
package tomita

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

// Parser это класс для работы с томитой
type Parser struct {
	execPath string
	confPath string
}

// New создает инстанс парсера
func New(execPath, confPath string) (Parser, error) {
	confDir := path.Dir(confPath)
	if confDir != "." {
		confPath = path.Base(confPath)
		os.Chdir(confDir)
	}
	if _, err := os.Stat(execPath); os.IsNotExist(err) { // TODO: check if it's executable
		return Parser{}, errors.New("the tomita path doesn't exist")
	}
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		return Parser{}, errors.New("the tomita path doesn't exist")
	}
	tp := Parser{
		execPath: execPath,
		confPath: confPath,
	}

	return tp, nil
}

// Run запускает парсер и считывает вывод
func (tp *Parser) Run(text string) (string, error) {
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

	_, err = stdin.Write([]byte(text))
	if err != nil {
		panic("Tomita probably hasn't waited for the input. It can be caused by an error in config.proto. " +
			err.Error())
	}
	stdin.Close()

	output, err := ioutil.ReadAll(stdout)
	panicOnErr(err)

	return string(output), nil
}

func panicOnErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
