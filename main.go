package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Hello World"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0])
	fmt.Printf("%c", s[0])
	fmt.Println()
	fmt.Println(s[0:6])
	fmt.Println(s[:6])
	fmt.Println(s[6:])
	fmt.Println()

	s = s + " Again"
	s += " Again"
	fmt.Println(s)
	fmt.Println()

	fmt.Println("Word \n Space")
	fmt.Println("Word \t Tab")
	fmt.Println("Word \bBackspace")

	abc := "abc"
	b := []byte(abc)
	fmt.Println(abc, b)
	fmt.Printf("%s, %s", abc, b)
	fmt.Println()
	abc2 := string(b)
	fmt.Println(abc2)

	fmt.Println(strings.ToTitle("Hello World"))
	fmt.Println(strings.ToLower("Hello World"))
	fmt.Println(strings.ToUpper("Hello World"))
	fmt.Println(strings.HasPrefix("Hello World", "Hello")) //true
	fmt.Println(strings.HasSuffix("Hello World", "Hello")) //false
	fmt.Println(strings.HasPrefix("Hello World", "World")) //false
	fmt.Println(strings.HasSuffix("Hello World", "World")) //true
	fmt.Println(strings.Replace("Hello World World", "World", "All", 1))
	fmt.Println(strings.Replace("Hello World World", "World", "All", 2))
	fmt.Println(strings.ReplaceAll("Hello World World", "World", "All"))

	//*** STRING BUILDER ***
	var sb strings.Builder
	fmt.Println("This a builder: ", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.WriteString("Hallo")
	fmt.Println("This a builder: ", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.Grow(100)
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println(sb.String())
	sb.Reset()
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println(sb.String())

	sb.Write([]byte{65, 65, 65})
	fmt.Println(sb.String())

	x := 123
	// Int to string
	y := strconv.Itoa(x)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T \n", y)

	// String to int
	z, _ := strconv.Atoi(y)
	fmt.Println(z)
	fmt.Printf("%T \n", z)
}
