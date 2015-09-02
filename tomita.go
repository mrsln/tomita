// Package tomita – это обертка для Томита-парсера
package tomita

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

// Parser это класс для работы с Томитой
type Parser struct {
	execPath   string
	confPath   string
	originalWd string
}

// Output является структурой ответа Томиты
type Output struct {
	Document document `xml:"document"`
}

type document struct {
	Facts facts  `xml:"facts"`
	Leads []Lead `xml:"Leads>Lead"`
}

type facts struct {
	Facts []Fact `xml:",any"`
}

// Fact содержит факт
type Fact struct {
	// Название фактов
	XMLName xml.Name
	// факты со значением
	Values []FactValue `xml:",any"`
}

// FactValue это конкретное значение из факта
type FactValue struct {
	// название факта
	XMLName xml.Name
	// спарсенное значение
	Value string `xml:"val,attr"`
}

// Lead это предложение содержащее факт
type Lead struct {
	Text string `xml:"text,attr"`
}

// Result очищенный ответ, после парсинга xml
type Result struct {
	Facts []Fact
	Leads []Lead
}

// New создает инстанс парсера
func New(execPath, confPath string) (Parser, error) {
	if _, err := os.Stat(execPath); os.IsNotExist(err) { // TODO: check if it's executable
		return Parser{}, errors.New("the tomita path doesn't exist: " + execPath)
	}
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		return Parser{}, errors.New("the config path doesn't exist: " + confPath)
	}
	wd, _ := os.Getwd()
	tp := Parser{
		execPath:   execPath,
		confPath:   confPath,
		originalWd: wd,
	}

	return tp, nil
}

// Run запускает парсер и считывает вывод
func (tp *Parser) Run(text string) (Result, error) {
	os.Chdir(path.Dir(tp.confPath))
	cmd := exec.Command(tp.execPath, path.Base(tp.confPath))

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return Result{}, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return Result{}, err
	}

	err = cmd.Start()
	if err != nil {
		return Result{}, err
	}

	_, err = stdin.Write([]byte(text))
	if err != nil {
		return Result{}, errors.New("Tomita probably hasn't waited for the input. " +
			"It can be caused by an error in config.proto. " +
			err.Error())
	}
	stdin.Close()

	output, err := ioutil.ReadAll(stdout)
	if err != nil {
		return Result{}, err
	}

	os.Chdir(tp.originalWd)

	var out Output
	err = xml.Unmarshal(output, &out)
	if err != nil {
		return Result{}, err
	}

	r := Result{Facts: out.Document.Facts.Facts, Leads: out.Document.Leads}

	return r, nil
}

func panicOnErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
