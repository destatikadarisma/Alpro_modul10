package main
import "fmt"
func main(){
	var a, b, durasi int
	fmt.Scan(&a, &b)
	if b>a {
		durasi = b-a
	} else if b<a {
		durasi = (24-a) + b
	} else {
		durasi = 0
	}
	fmt.Println(durasi)

}