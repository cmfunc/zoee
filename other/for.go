package other

/*
for 循环中v是
*/

func Forf()  {
	m:=make(map[int]string, 0)
	s:=[3]string{"A","B","C"}
	for i, v := range s {
		println(&v)
		m[i]=v
	}
}