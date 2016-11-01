//Program to Print Command Line Arguments in a Single Line
package main
import (
  "fmt"
  "os"
)
func main (){
  var s1,s2 string
  var i int
  for i=1;i<len(os.Args); i++{
      s1+=s2+os.Args[i]
      s2=""
  }
  fmt.Printf("%s \n",s1)
}
