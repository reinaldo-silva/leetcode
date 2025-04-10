
# Sliding Window (Janela Deslizante)

## 📖 O que é Sliding Window?

Sliding Window é uma técnica de resolução de problemas amplamente utilizada em algoritmos, especialmente para problemas que envolvem arrays ou listas. A ideia principal é manter uma "janela" de elementos contíguos que se desloca através da estrutura de dados, permitindo que você processe ou analise um subconjunto dos dados de maneira eficiente.

Essa técnica é particularmente útil para problemas que requerem a busca de subarrays ou subsequências que atendam a determinadas condições, como a soma máxima, média, ou a presença de duplicatas.

## 🔍 Como Funciona?

A técnica de Sliding Window pode ser dividida em duas abordagens principais:

1. **Janela de Tamanho Fixo**: A janela tem um tamanho fixo e se desloca pela estrutura de dados.
2. **Janela de Tamanho Variável**: A janela pode crescer ou encolher dinamicamente, dependendo das condições do problema.

## 📊 Exemplos

### Exemplo 1: Máxima Soma de Subarray de Tamanho K

**Problema**: Dado um array de inteiros e um inteiro `K`, encontre a soma máxima de qualquer subarray contíguo de tamanho `K`.

**Solução em Go**:

```go
package main

import (
    "fmt"
    "math"
)

func maxSumSubarray(arr []int, K int) int {
    maxSum := math.MinInt64
    windowSum := 0

    for i := 0; i < K; i++ {
        windowSum += arr[i]
    }

    for i := K; i < len(arr); i++ {
        windowSum = windowSum - arr[i-K] + arr[i]
        maxSum = max(maxSum, windowSum)
    }

    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    arr := []int{2, 1, 5, 1, 3, 2}
    K := 3
    fmt.Println("Máxima Soma de Subarray de Tamanho K:", maxSumSubarray(arr, K))
}
```

**Complexidade de Tempo**: O(n), onde n é o número de elementos no array. A abordagem percorre o array uma única vez.

### Exemplo 2: Longest Substring Without Repeating Characters

**Problema**: Dada uma string, encontre o comprimento da substring mais longa que não contém caracteres repetidos.

**Solução em Go**:

```go
package main

import (
    "fmt"
)

func lengthOfLongestSubstring(s string) int {
    charSet := make(map[rune]bool)
    left, maxLength := 0, 0

    for right, char := range s {
        for charSet[char] {
            delete(charSet, rune(s[left]))
            left++
        }
        charSet[char] = true
        maxLength = max(maxLength, right-left+1)
    }

    return maxLength
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    s := "abcabcbb"
    fmt.Println("Comprimento da Longest Substring Without Repeating Characters:", lengthOfLongestSubstring(s))
}
```

**Complexidade de Tempo**: O(n), onde n é o comprimento da string. A abordagem percorre cada caractere da string uma vez.

## 🔑 Vantagens do Sliding Window

- **Eficiência**: Reduz a complexidade de tempo de O(n^2) para O(n) na maioria dos casos, evitando iterações desnecessárias.
- **Simplicidade**: A técnica é relativamente fácil de entender e implementar, especialmente para problemas que envolvem sequências contíguas.

## 📌 Conclusão

A técnica de Sliding Window é uma poderosa ferramenta para resolver problemas de algoritmos, especialmente aqueles que envolvem arrays e strings. Ao entender como aplicar essa técnica, você pode otimizar sua abordagem e melhorar a eficiência de suas soluções.

Se você tiver dúvidas ou quiser discutir mais sobre essa técnica, sinta-se à vontade para entrar em contato!

