package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main()  {
	var a string = "abc"
	b := strings.HasPrefix(a, "a")
	c := strings.HasSuffix(a, "c")
	d := strings.Contains(a, "bc")
	e := strings.Index(a, "c")
	f := strings.Count(a, "a")
	g := strings.LastIndex(a, "a")
	o := strings.Replace(a, "a", "A", -1)
	m := strings.Repeat(a, 3)
	n := strings.ToLower(a)
	p := strings.ToUpper(a)
	q := strings.Trim(a, "t")
	r := strings.TrimSpace(a)
	s := strings.Fields(a)
	t := strings.Split(a, "b")
	u := strings.Join(a, ':') // a is slice
	v := strings.NewReader(a)
	w := strconv.Atoi(a)
	y := strconv.FormatInt(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
