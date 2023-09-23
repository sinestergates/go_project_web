package main
import (

	"fmt"
	"time"
)


func main() {
	i:= 10
	go fmt.Printf("1. Значение переменной i равно %d\n", i)
	i++
	go fmt.Printf("2. Значение переменной i равно %d\n", i)
	go func() {
	i++
	go fmt.Printf("3. Значение переменной i равно %d\n", i)
	}()
	i++
	go fmt.Printf("4. Значение переменной i равно %d\n", i)
	time.Sleep(1000000)
}