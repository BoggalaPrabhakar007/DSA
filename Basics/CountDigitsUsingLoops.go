import "fmt"

func main() {
	DigitCount(1234)	
}

func DigitCount(num int) {
	count := 0
	for num != 0 {
		num = num / 10
		count++
	}
	fmt.Println("Digit Count:", count)
}
