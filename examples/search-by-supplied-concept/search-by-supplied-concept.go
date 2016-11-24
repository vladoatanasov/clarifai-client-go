package main

import (
	cl "github.com/mpmlj/clarifai-client-go"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("..")
	viper.SetConfigName("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	sess := cl.NewSession(viper.GetString("clarifai_api.client_id"), viper.GetString("clarifai_api.client_secret"))

	err = sess.Connect()
	if err != nil {
		panic(err)
	}

	// Search By User Supplied Concept.
	svc := cl.NewSearchService(sess)

	resp, err := svc.SearchByUserSuppliedConcepts([]string{"album"})
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}
