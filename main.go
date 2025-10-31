package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Считываем операцию
	var op string
	fmt.Println("Введите тип операции (AVG - среднее, SUM - сумму, MED - медиану)")
	fmt.Scan(&op)

	// Считываем строку с числами через запятую
	fmt.Println("Введите числа через запятую (например, 2, 10, 9)")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	// Преобразуем строку в срез int
	slice := parseNumbers(s)

	// В зависимости от операции вызываем нужную функцию
	switch op {
	case "AVG":
		avgOperation(slice)
	case "SUM":
		sumOperation(slice)
	case "MED":
		medOperation(slice)
	default:
		fmt.Println("Неизвестная операция")
	}
}

// parseNumbers разбивает строку по запятой и преобразует в []int
func parseNumbers(s string) []int {
	parts := strings.Split(s, ",")
	var nums []int
	for _, part := range parts {
		part = strings.TrimSpace(part)
		n, err := strconv.Atoi(part)
		if err == nil {
			nums = append(nums, n)
		} else {
			fmt.Println("Ошибка в числе:", part)
		}
	}
	return nums
}

func avgOperation(slice []int) {
	if len(slice) == 0 {
		fmt.Println("Нет чисел для вычисления.")
		return
	}
	sum := 0
	for _, v := range slice {
		sum += v
	}
	avg := float64(sum) / float64(len(slice))
	fmt.Printf("Среднее значение: %.2f\n", avg)
}

func sumOperation(slice []int) {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	fmt.Printf("Сумма: %d\n", sum)
}

func medOperation(slice []int) {
	n := len(slice)
	if n == 0 {
		fmt.Println("Нет чисел для вычисления.")
		return
	}

	sorted := make([]int, n)
	copy(sorted, slice)
	sort.Ints(sorted)

	if n%2 == 1 {
		median := sorted[n/2]
		fmt.Printf("Медиана: %d\n", median)
	} else {
		median := float64(sorted[n/2-1]+sorted[n/2]) / 2
		fmt.Printf("Медиана: %.2f\n", median)
	}
}

//stop
