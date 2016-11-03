package main
import(
  "fmt"
  "os"
  "io/ioutil"
  "strings"
)

func main(){
count:=make(map[string]int)
files:=os.Args[1:]

for _,filename:=range(files){
      data,err:=ioutil.ReadFile(filename)
        if err !=nil {
          fmt.Fprintf(os.Stderr,"%v",err)
          continue
      }

for _,line:=range strings.Split(string(data),"\n"){
    count[line]++
}

for line,number:= range count{
    if number>1 {
      fmt.Printf("%d \t %s\n",number,line)
    }
}
}
}
