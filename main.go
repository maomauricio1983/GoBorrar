package main

import "fmt"

/*
// funciones
func main() {
	miFuncion()
	miFuncionConParametros(2, 6)
	fmt.Print(miFuncionConRetorno("Mauricio"))
	nombre, apellido, edad := miFuncionConRetornoMultiple()
	fmt.Print("Hola %d %d, tu edad es %d años", nombre, apellido, edad)

	// funcion anonima
	println("lA SUMA ES IGUAL A :  ", suma(5, 8))

	Tabla := tabla(2)
	for i := 1; i < 10; i++ {
		fmt.Printf("2 x %v = %v ", i, Tabla())
	}
}

func miFuncion() {
	fmt.Print("Hola mundo")
}

func miFuncionConParametros(n1 int, n2 int) {
	resultado := n1 + n2
	fmt.Print(resultado)
}

// funcion con retorno
func miFuncionConRetorno(nombre string) string { // le indico que retorna un string
	return "tu nombre es: " + nombre
}

func miFuncionConRetornoMultiple() (string, string, int) {
	return "Juan", "Perez", 41
}

// funcion anonima
var suma = func(numero1 int, numero2 int) int {
	return numero1 + numero2
}

// closure ( funciones que devuelven funciones)
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero + secuencia
	}
}
*/

/*
// Go rutines y channels
func main() {
	fmt.Print(miFuncion("mauricio"))
	time.Sleep(time.Second * 5)
	fmt.Println(miFuncion("maria"))

	// ejemplo 2
	miCanal := make(chan string)
	go func() {
		miCanal <- miFuncion("Pedro")
	}()
	fmt.Println(<-miCanal)
	fmt.Println("Continuar con la ejecucion")
}

func miFuncion(parametro string) string {
	return "hola " + parametro
}

*/

func main() {
	miFuncion()
}

func miFuncion() {
	defer fmt.Println("este es el mensaje final")
	fmt.Println("este es el primer mensaje")
}
