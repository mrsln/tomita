Go Tomita Parser wrapper
========================

Небольшой враппер для удобной работы с Томита Парсер от Яндекса в Go.

Пример
------

```go
import "github.com/mrsln/tomita"

func main() {
	p, err := tomita.New("/bin/tomita", "example/config.proto")
	...
	output, err := p.Run("This is text to parse")
	...
}

```

Важно
-----

- config.proto не должен содержать дескрипторов File (ввод/вывод осуществляется через STDIN/STDOUT)


За пример спасибо автору [poor-python-yandex-tomita-parser](https://github.com/vas3k/poor-python-yandex-tomita-parser).
