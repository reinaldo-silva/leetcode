# 🧠 Análise Big O do Algoritmo Two Sum em Go

## 📌 Código analisado

```go
func twoSum(nums []int, target int) []int {
	hasher := make(map[int]int)

	for index, num := range nums {
		if val, found := hasher[num]; found {
			return []int{val, index}
		}
		hasher[target-num] = index
	}

	return []int{}
}
```

---

## 📊 Complexidade de Tempo — Big O (Time Complexity)

### 1. `make(map[int]int)`
- Cria uma hash table vazia.
- Operação constante.
- **Complexidade:** `O(1)`

---

### 2. `for index, num := range nums`
- Percorre todos os elementos do array `nums`.
- `n` é o número de elementos em `nums`.
- **Complexidade:** `O(n)`

---

### 3. `if val, found := hasher[num]; found`
- A verificação e leitura em um `map` é feita em tempo constante na média (implementado como uma hash table).
- **Complexidade:** `O(1)` por iteração.

---

### 4. `hasher[target - num] = index`
- A inserção em uma hash table (`map`) também ocorre em tempo constante na média.
- **Complexidade:** `O(1)` por iteração.

---

### 🔁 Tudo isso está dentro do `for`, ou seja, acontece `n` vezes:

- Operações constantes (`O(1)`) dentro de um loop de `n` itens.
- **Complexidade total de tempo:** `O(n)`

---

## 🧠 Complexidade de Espaço — Big O (Space Complexity)

### Estrutura adicional usada:
- `hasher`, o map utilizado para armazenar valores já vistos.
- No pior caso (nenhum par encontrado até o final), o `map` armazena todos os `n` elementos do array.

### Portanto:
- **Complexidade de espaço:** `O(n)`

---

## ✅ Resumo

| Tipo             | Complexidade |
|------------------|--------------|
| Tempo (Time)     | `O(n)`       |
| Espaço (Memory)  | `O(n)`       |

---

## 🧪 Exemplo de execução

```go
nums := []int{2, 7, 11, 15}
target := 9
```

### Iterações:
1. num = 2 → salva `7:0` no map.
2. num = 7 → encontra no map → retorna `[0, 1]`.

Pouquíssimas operações — muito eficiente! 🚀

---

## 🆚 Comparação com abordagem O(n²)

Uma abordagem mais simples (mas lenta) envolveria dois `for` aninhados, o que resultaria em `O(n²)`:

```go
for i := 0; i < len(nums); i++ {
	for j := i + 1; j < len(nums); j++ {
		if nums[i] + nums[j] == target {
			return []int{i, j}
		}
	}
}
```

---

## 🧠 Conclusão

- A solução com `map` (hash table) reduz o tempo de execução de `O(n²)` para `O(n)`.
- Essa é a melhor abordagem possível para esse problema usando estrutura de dados básica.
