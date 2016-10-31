package main
import "fmt"
func main(){
fmt.Println("Enter the number")
var j int
fmt.Scan(&j)
fmt.Println("power of 2 true:",check_power(j))
}
func check_power(x int)(y bool){
for((x%2)==0){
if x>1{
x=x/2
}
}
if (x==1){
y=true
return y
}
y=false
return y
}
