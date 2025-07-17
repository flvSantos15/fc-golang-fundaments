package ponteiros

import "fmt"

func Test() {
	// ponteiros
	// memoria endereco(40) <---- a <----- 40

	a := 10
	
	var ponteiro *int = &a
	fmt.Println(*ponteiro)

	*ponteiro = 50
	fmt.Println(a)
	fmt.Println(*ponteiro)

	variavel := 10

	abc(&variavel)
}

func abc(a *int) {
	*a = 200
}