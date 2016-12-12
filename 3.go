package main
import(
  "fmt"
)
func largest_prime_factor (limit int) int{
  i :=2
  for i*i < limit{
    for limit % i == 0{
      limit = limit/i
    }
    i+=1
  }
  return limit
}
func main(){
  fmt.Println(largest_prime_factor(600851475143))
}
