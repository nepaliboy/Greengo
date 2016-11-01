//simple infinite loop equivalent to while statement
//for is the only loop statement in the Go
package main
import(
  "fmt"
)
func main() {
var x int
x=1
for x>0 {
  fmt.Printf("i am infinite because of %v greater than 0 \n",x)
}
}
