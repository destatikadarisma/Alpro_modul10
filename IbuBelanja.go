package main
import "fmt"
func main(){
	var daging, masakan string
	fmt.Scan(&daging,&masakan)
	if daging == "ayam" {
		fmt.Println(" Daging : Ayam")
		if masakan == "sop" {
			fmt.Println(" tambahan : sayur")
			} else if masakan == "katsu"  {
				fmt.Println(" tambahan : tepung")
			}
	} else if daging == "sapi" {
		fmt.Println(" Daging : Sapi")
		if masakan == "teriyaki"{
			fmt.Println(" tambahan : bombay")
		}else if masakan == "semur"{
			fmt.Println("tambahan : kecap ")

		}
	}
	
}