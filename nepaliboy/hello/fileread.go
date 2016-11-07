package main
import(
  "fmt"
  "io/ioutil"
  "os"
  )
func main(){
  fname:=os.Args[1]
  fmt.Printf(" Reading File.....%s \n",fname)
  input,err:=ioutil.ReadFile(fname)
  if err!=nil {
    fmt.Printf("%v",err)
  }
  fmt.Printf("%v",input)
}
