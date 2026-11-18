package main
import "fmt"
func main(){
	var s int
	fmt.Scan(&s)
	if s >= 310 && s <= 677 {
		if s >= 450 {
			fmt.Println("KELAS REGULER")
		} else {
			fmt.Println("KELAS MATRIKULASI")
		}
	}
}