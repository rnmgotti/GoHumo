package taskd4

import (
	"fmt"
)
func Table() {

	fmt.Println("Вы выбрали задание Вывод таблицы умножения.")
	for i := 1; i < 10; i++ {
	fmt.Println()
    for j := 1; j < 10; j++ {
	fmt.Println(i, "x", j, "=", i*j)		
       }
	}
}