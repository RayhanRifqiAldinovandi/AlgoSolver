package main


type Stack struct{
	items []string
}


//Use uppercase so function can be accesed from another file	
func (s* Stack) Pop() string{
	if len(s.items) == 0{
		return "Stack is empty"
	}
	//get the top element
	item := s.items[len(s.items)-1]
	//create new slice but leave out the top element
	s.items = s.items[:len(s.items)-1]

	return item
}

func (s* Stack) Push(character string){
	s.items = append(s.items, character)
}

func (s* Stack) Size() int{
	return len(s.items)
}

func (s* Stack) Top() string{
	if len(s.items) == 0{
		return "Stack is empty"
	}
	//get the top element
	item := s.items[len(s.items)-1]

	return item
}