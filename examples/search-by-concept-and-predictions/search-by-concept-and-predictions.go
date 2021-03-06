package main

import (
	"os"

	cl "github.com/mpmlj/clarifai-client-go"
)

func main() {

	sess, err := cl.Connect(os.Getenv("CLARIFAI_API_ID"), os.Getenv("CLARIFAI_API_SECRET"))
	if err != nil {
		panic(err)
	}

	q := cl.NewAndSearchQuery()
	q.WithUserConcept("album") // inputs
	//q.WithoutUserConcept("vacation") // outputs
	//q.WithAPIConcept("singer") // outputs
	//q.WithoutAPIConcept("singer") // outputs

	resp, err := sess.Search(q).WithPagination(1, 5).Do()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)

}
