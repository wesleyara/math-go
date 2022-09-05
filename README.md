# Math Go

O Math Go é um projeto de estudo de Go, que tem como objetivo implementar funções matemáticas. Com ele estou aprendendo mais sobre a linguagem e disponibilizando o código para quem quiser estudar também.

## Como usar

Para usar o Math Go, basta chamar as funções contidas no módulo "utils" e passar os parâmetros necessários. Por exemplo, para calcular o seu IMC, basta chamar a função "math.IMC()" e passar o número como parâmetro.

```go
package main

import (
  "fmt"
  "utils/math"
)

func main() {
  fmt.Println(math.IMC(70, 1.70))
}
```

## Funções

- Divide
- Multiply
- Sum
- Sub
- Div
- IMC

## Licença
Esse projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
