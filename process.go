package main
import (
	"fmt"
    "os"
	"io"
	"strconv"
	"log"
	"io/ioutil"
	"strings"
)

func lab3(path_from string, path_to string, file_name string) {
	






    file, err := os.Open(path_from + "/" + file_name) //open
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close() 
     
    data := make([]byte, 2000)
	count := 0
	str :="";
    for{
		
        n, err := file.Read(data)
        if err == io.EOF{   // если конец файла
            break           // выходим из цикла
		}
		str += string(data[:n])
        //fmt.Print(string(data[:n]))
	}
	
	for i := 0; i < len(str); i++ {
		if (string(str[i]) == " ") && (string(str[i - 1]) == "." || string(str[i - 1]) == "!" || string(str[i - 1]) == "?"){
			count++
		}
		if ((i + 1) == len(str)) && (string(str[i]) == "." || string(str[i]) == "!" || string(str[i]) == "?") {
			count++
		}
	}
	text := strconv.Itoa(count)
	file_to, err := os.Create(path_to + "/" + new_file_name) //write
     
    if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1) 
    }
	defer file_to.Close() 
    file_to.WriteString(text)
}
 
func main() {
	
	fmt.Print("Total number of processed files: " + strconv.Itoa(i))
	//time.Sleep(1000)
	fmt.Scanf(&input)



}