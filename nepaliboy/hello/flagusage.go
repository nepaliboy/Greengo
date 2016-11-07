package main
import(
  "fmt"
  //"strings"
  "flag"
)

func main(){
  usage()
  fmt.Printf("Testing the usage of Flag \n ")
  fmt.Printf("other arguments: \n",flag.Args())
}

func usage(){
  var n=flag.Bool("n",true, "Add 2 Trailing new line")
  var sep=flag.String("s","::","seperator")
  flag.Parse()
  fmt.Printf("%p %p",*n,*sep)
  //fmt.Print(strings.Join(flag.Args(), *sep))
  if *n {
    fmt.Printf("\n\n")
}

}
