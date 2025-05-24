package main
 import (
 "fmt"
 )

  type Nodo struct {
 valor int
 nextNodo *Nodo
 }

 type Pila struct {
 tope *Nodo
  size int 
 }

 func (p *Pila) push(element int) {
 p.tope = &Nodo{
 valor: element,
 nextNodo: p.tope,
 }
 p.size++
 }

  func (p *Pila) pop() int {
 if p.tope == nil {
	fmt.Println("Error: la pila está vacía.")

 return -1
 }

 valor := p.tope.valor
 p.tope = p.tope.nextNodo
 p.size--
return valor
}

 func (p *Pila) getTop() int {
 if p.tope == nil {
 fmt.Println("La pila está vacía.")
 return 0
 }
 return p.tope.valor
 }

 func (p *Pila) getSize() int {
 return p.size
 }

  func main() {

 p := &Pila{} // Creamos una pila vacía


 fmt.Println("Pila recién creada:")

 fmt.Println("Tamaño:", p.getSize())

 fmt.Println("Tope:", p.getTop()) // Verifica comportamiento con pila vacía

 p.push(10)
 p.push(20)
 p.push(30)

 fmt.Println("ingresamos 10, 20 y 30")

 fmt.Println("Tamaño actual:", p.getSize())
 fmt.Println("Elemento en el tope:", p.getTop())

 p.pop()

 fmt.Println("eliminamos el 30 mediante un pop")

 fmt.Println("Nuevo tope:", p.getTop())
 fmt.Println("Tamaño:", p.getSize())

 p.pop()
 
 

 fmt.Println("hacemos 1 pop")
  fmt.Println("Tope actual:", p.getTop())     // 
fmt.Println("Tamaño final:", p.getSize())
}

