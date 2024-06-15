package Stack

func Run() {
	s := New()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	s.Pop()
	s.Pop()
	s.Print()
	// fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())
}
