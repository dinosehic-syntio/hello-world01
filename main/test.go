package main
import "fmt"
func main() {
	fmt.Println("Hello world!")
	fmt.Print("Hello ")
	fmt.Print("world!\n")
	for i:=0; i<5; i++ {
		fmt.Println(i+1   )
	}

	n := "nesto"
	i := 0
	fmt.Println(n)
	for i < len(n){
		fmt.Printf("%c", n[i])
		i++
	}
	fmt.Println("Upisite prvi broj")
	var x int
	fmt.Scanf("%d", &x)
	fmt.Println("Upisite drugi broj")
	var y int
	fmt.Scanf("%d", &y)


	fmt.Println("Funkcija zbrajanja za ", x, " i ", y, ": ", add(x,y))


}

func add(a int, b int) int{
	sum := a + b
	return sum
}
