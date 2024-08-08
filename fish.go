package fish

import "github.com/Ajaykumar-N-GITHUB/marine"

func Fishing(s string) string {
	p1 := marine.Seafisher(s)
	return p1 + " is a fisher man"
}

func NotFishing(r string) string {

	p2 := marine.Teaching(r)
	return p2 + " not a fisher man"
}
