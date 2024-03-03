package taskd2

import (
	"fmt"
	"log"
	"math"
)
func Anniversaryyear() {

	log.Println("Программа запущена. Приветствую, я помогу вам узнать сколько вам осталось до юбилея!")

	var age int

	fmt.Println("Сколько вам лет?")
	fmt.Scanf("%d", &age)

	var anniversary int

	anniversary = int(math.Ceil(float64(age)/10))*10 - age

	fmt.Printf("До вашего юбилея осталось %d лет.\n", anniversary)
	log.Println("Программа завершена.")

}

func Areap() {

	var a int
	var b int
  
	fmt.Println("Чему равна длина первой стороны прямоугольника?")
	fmt.Scan(&a)
	fmt.Println("Чему равна длина второй стороны прямоугольника?")
	fmt.Scan(&b)
	fmt.Println("Площадь прямоугольника равен=" + fmt.Sprint(a*b))
  
}

func Perimeterk() {

	var a int
	
	  fmt.Println("Чему равна сторона квадрата?")
	  fmt.Scan(&a)
	  fmt.Println("Периметр квадрата равен=" + fmt.Sprint(4*a) + "!")
	
}

func Perimetrp() {

	var a int
	var b int
	
	  fmt.Println("Чему равна длина первой стороны прямоугольника?")
	  fmt.Scan(&a)
	  fmt.Println("Чему равна длина второй стороны прямоугольника?")
	  fmt.Scan(&b)
	  fmt.Println("Периметр прямоугольника равен=" + fmt.Sprint(2*(a+b)))
	
}