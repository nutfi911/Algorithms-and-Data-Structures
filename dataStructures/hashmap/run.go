package hashmap

import "fmt"

func Run() {

	ht := New(4)

	fmt.Println(ht.Size)
	fmt.Println(ht.Set("cyan", "4"))
	fmt.Println(ht.Set("white", "#FFFFF"))
	fmt.Println(ht.Set("whitey", "#FFFFF"))
	fmt.Println(ht.Set("cyn", "422"))
	fmt.Println(ht.Set("1", "422"))
	fmt.Println(ht.Set("123123321fekwobjgfrjwefiijwf", "422"))
	fmt.Println(ht.KeyMap)
	fmt.Println("Get cyn", ht.Get("cyn"))
	fmt.Println("Get all keys", ht.Keys())
	fmt.Println("Get all values", ht.Values())
}
