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

	// Search By Predicted Concept.
	q := cl.NewAndSearchQuery()
	q.WithAPIConcept("vacation") // inputs

	resp, err := sess.Search(q).Do()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
