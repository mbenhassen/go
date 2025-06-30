package main

import (
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
   dat,err := os.ReadFile(filename)
   if err != nil {
	  return "", err
   }
   if len(dat) == 0 {	  
	 //return "", errors.New("le fichier est vide")
	  return "", fmt.Errorf("le fichier %s est vide", filename)
   }
   return string(dat), nil
}

func main() {
	dat, err := ReadFile("fichier.txt")
	if err != nil { 
		fmt.Printf("Erreur lors de la lecture du fichier: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Contenu du fichier: %s\n", dat)
}
