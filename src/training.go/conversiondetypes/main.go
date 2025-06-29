package main
import "fmt"

func main() {
   var pourcentage float64 = 85.0

   pourcentage = pourcentage + 2
   pourcentage += 1
   fmt.Printf("Le pourcentage est %f%%\n", pourcentage)
   fmt.Printf("La valeur est %d\n", int(pourcentage))

   n := 42
   f := float64(n) + 0.5
   fmt.Printf("La valeur de n est %d et de f est %f\n", n, f)


}
