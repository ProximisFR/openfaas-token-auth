package openfaastokenauth

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Auth Compare provided token with function-auth-token (Kubernetes Secret) to allow OpenFaaS function to execute or not
func Auth() {
	token, err := ioutil.ReadFile("/var/openfaas/secrets/function-auth-token")
	if err != nil {
		err.Error()
		os.Exit(1)
	}

	if string(token) == string(os.Getenv("Http_Function_Auth_Token")) {
		return
	}
	fmt.Print("You're not authenticated at function level. Exit.")
	os.Exit(1)
}
