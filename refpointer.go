package main

import "github.com/davecgh/go-spew/spew"

func main() {
	s := 1

	r := &s
	rr := &r
	rrr := &rr

	vvv := *rrr
	vv := *vvv
	v := *vv

	spew.Dump(s, r, rr, rrr, &rrr, vvv, vv, v)

}
