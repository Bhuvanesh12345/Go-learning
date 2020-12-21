package main

import "fmt"

func foostr1() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	for i := 0; i < len(sample); i++ {
		slice := sample[i:len(sample)]
		fmt.Println("Println:")
		fmt.Printf("%x\n", slice)
		fmt.Println("Printf with % q:")
		fmt.Printf("%q\n", slice)

		fmt.Println("Printf with % +q:")
		fmt.Printf("%+q\n", slice)
		fmt.Println("==================================")
	}
}

func main() {
	//foostr1()
	//foostr2()
	rngloop()
	udemy_str()
}

func foostr2() {
	const placeOfInterest = `⌘`
	fmt.Printf("plain string: ")
	fmt.Printf("%T", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}

func rngloop() {
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}

func udemy_str() {
	s := "Hello, playground"
	fmt.Println(s)
	s = "Hello, 世界"
	s = "��=� ⌘"

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(bs); i++ {
		fmt.Printf("%#U ", bs[i])
	}
	fmt.Printf("\n")
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
	fmt.Printf("\n")
	for i, v := range bs {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
