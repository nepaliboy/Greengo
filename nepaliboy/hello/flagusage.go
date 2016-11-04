package main
import(
  "fmt"
  "strings"
  "flag"
)

func main(){
  usage()
  fmt.Printf("Testing the usage of Flag \n ")
  fmt.Printf("other arguments: \n",flag.Args())
}

func usage(){
  var n=flag.Bool("n",false, "Omit Trailing new line")
  var sep=flag.String("s","::","seperator")
  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), *sep))
  if !*n {
    fmt.Println()
  }

}
