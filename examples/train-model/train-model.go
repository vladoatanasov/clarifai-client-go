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

	resp, err := sess.TrainModel("music-model-id-1").Do()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
