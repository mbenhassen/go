package main

import (
	"fmt"
	"strings"
)

type Post struct {
	Title   string
	Text    string
	published bool
}

func (p Post) Headline() string {
	return fmt.Sprintf("Post: %v - %v", p.Title, p.Text[:50])

}

func (p Post) Published() bool{
	return p.published
}

func (p *Post) Publish(){
	p.published = true
}

func (p *Post) Unpublish(){
	p.published = false
}

func (p *Post) UpperTitle() {
	p.Title = strings.ToUpper(p.Title)
}


func main() {
	p := Post{
		Title:   "Go Programming",
		Text:    `An introduction to 
		Go programming language
		KLFELRHFEJRHLER
		KJGLKSLKJhgehgjt
		`,
}
	fmt.Printf("Title:%v\n", p)
	fmt.Println(p.Headline())
	fmt.Printf("Published: %v\n", p.Published())
	p.Publish()
	fmt.Printf("Published: %v\n", p.Published())
	p.Unpublish()
	fmt.Printf("Published: %v\n", p.Published())
	p.UpperTitle()
	fmt.Printf("Upper Title: %v\n", p.Title)

	pythonPost := &Post{
		Title:   "Python Programming",
		Text:    `Python (prononcé /pi.tɔ̃/) est un langage de 
		programmation interprété, multiparadigme et multiplateformes. Il favorise la programmation impérative structurée, fonctionnelle et orientée objet. Il est doté d'un typage dynamique fort, d'une gestion automatique de la mémoire par ramasse-miettes 
		et d'un système de gestion d'exceptions ; il est ainsi similaire à Perl, Ruby, Scheme, Smalltalk et Tcl.`,
}
	pythonPost.UpperTitle()
	fmt.Printf("Title:%v\n", pythonPost.Headline())
}