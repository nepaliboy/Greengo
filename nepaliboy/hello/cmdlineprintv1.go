package main
import (
"fmt"
"os"
)
func main() {
fmt.Println(" All Arguments are stored in Slice ",os.Args)
var i int
for i=0;i<len(os.Args);i++{
fmt.Printf("the %v argument is %v \n",i,os.Args[i])
}
}
