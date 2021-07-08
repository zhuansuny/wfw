package wfw

import (
	"testing"
)


func Test_Nodeinsert(t *testing.T) {
	cases := []struct{ input string }{
		{"/api/login"},
		{"/api/login/user"},
		{"/api/logout"},
		{"/api/db/qurey"},
		{"/api/db/qurey/list"},
		{"/api/db/qurey/*"},
	}
	n := node{}

	for _, cas := range cases {
		parts := parsePattern(cas.input)
		n.insert(cas.input, parts, 0)

		res := n.search(parts, 0)
		if res == nil || res.pattern != cas.input {
			t.Fatal("fail")
		}
	}

}
