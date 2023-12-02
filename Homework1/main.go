package main

const chances = 3

func main() {

	/* 1- misol

	son := 123

	yuzlar := son / 100
	oNlar := (son % 100) / 10
	birlar := son % 10

	fmt.Printf(" %d%d%d\n", yuzlar, birlar, oNlar)
	*/

	/* 2- misol

	son := (1245 % 1000) / 100

	fmt.Printf("yuzlar %d\n", son)
	*/

	/* 3- misol

	son := 1200 / 1000

	fmt.Printf("minglik %d\n", son)

	*/

	/* 4- misol

	sekund := 3220

	minut := sekund / 60

	fmt.Printf("Minut %d\n", minut)

	*/

	/* 5- misol

	sekund := 3820

	soat := sekund / 3600

	fmt.Printf("Soat %d\n", soat)

	*/

	/* 6- misol

	sekund := 3920

	minut := (sekund / 60) % 60

	soat := sekund / 3600

	fmt.Printf("%d Soat %d minut \n", soat, minut)

	*/

	/* 7- misol

	sekund := 3920

	sekund2 := sekund % 3600

	soat := sekund / 3600

	fmt.Printf("%d Soat %d sekund \n", soat, sekund2)

	*/

	/* 8- misol

	sekund := 3920

	sekund2 := (sekund % 3600) % 60

	minut := (sekund / 60) % 60

	soat := sekund / 3600

	fmt.Printf("%d Soat %d minut %d sekund \n", soat, minut, sekund2)

	*/

}
