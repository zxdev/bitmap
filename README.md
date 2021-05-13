# bitmap

A dynamic bitmap array that supports building a bitmap from programatically defined constant values OR build a list from a resource such as a text file for file directory. When the order is not known or may change passing in a map[string]int pointer will populate and can function as a helper lookup tool.

example use case

```golang

const (
	Read = 1 << iota
	Write
	Update
)

func main() {

  const 
	b := NewBitMap(nil, nil)
	b.Set(Write)
	b.Set(5)
	b.Set(66)
	fmt.Printf("%b %t %t\n", b, b.IsSet(Write), b.IsSet(63))
  // &[100100 100] true false
}

```
