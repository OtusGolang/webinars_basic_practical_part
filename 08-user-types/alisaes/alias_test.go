package alisaes

import (
	"fmt"
	"testing"
)

// Псевдоним типа int (alias)
type MyInt = int

// Новый пользовательский тип на основе int
type YourInt int

func TestAliases(t *testing.T) {
	var a int = 10
	var b MyInt = a            // Псевдоним можно использовать напрямую
	var c YourInt = YourInt(a) // А тут нужно явное преобразование
	// var cc YourInt = a      // Ошибка: нельзя присвоить int к YourInt напрямую

	fmt.Println(a, b, c)
}
