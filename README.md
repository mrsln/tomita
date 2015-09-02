Tomita-parser wrapper in Go
========================

Обертка [Томита-парсера](https://tech.yandex.ru/tomita/) в Go.

Пример
------

```go

p, err := tomita.New("/bin/tomita", "example/config.proto")
output, err := p.Run("This is a text to parse")

```

Важно
-----

- config.proto не должен содержать дескрипторов File (ввод/вывод осуществляется через STDIN/STDOUT)
