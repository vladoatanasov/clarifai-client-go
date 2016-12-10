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

	data := cl.InitInputs()

	// Option A. Adding an image from URL.
	_ = data.AddImageInput(cl.NewImageFromURL("https://samples.clarifai.com/metro-north.jpg"), "")

	// Option B. Adding an image from a local file.
	// NOTE. Currently API does not accept a mix of URL and base64 - based images!
	//i, err := cl.NewImageFromFile("../Dave_Gahan_New_York_2015-10-22.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//_ = data.AddImageInput(i, "")

	// As per https://developer-preview.clarifai.com/guide/predict#predictBy,
	// general model is used by default (ID "aaa03c23b3724a16a56b629203edc62c"),
	// but you can also set your own model using a SetModel() method on your input.
	data.SetModel(cl.PublicModelTravel)

	resp, err := sess.Predict(data).Do()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
