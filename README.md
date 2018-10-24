# OpenFaaS Auth Token

Simple helper (insecure) to auth with a token in header when calling OpenFaaS function and compare it with Kubernetes Secret named `function-auth-token` with `function-auth-token` as key for token.

## Usage

Create your Kubernetes secret :

```bash
kubectl create secret generic function-auth-token --from-literal=function-auth-token=YOURMARVELLOUSTOKEN
```

Call your function with token header :

```bash
echo "" |faas invoke YOURFUNCTION -H "Function-Auth-Token=YOURMARVELLOUSTOKEN"
```

```
curl --silent -H "Function-Auth-Token: YOURMARVELLOUSTOKEN" http://YOUROPENFAASGATEWAY/function/YOURFUNCTION
```

Checked environment variable (OpenFaaS put request headers into environment variables) is `Http_Function_Auth_Token`.

If token is different from Kubernetes Secret it will launch Golang panic func.

```golang
package main

import (
	"time"

	openfaasauthtoken "gitlab.omn.proximis.com/golang/public/openfaas-auth-token"
)

func main() {
    openfaasauthtoken.Auth()
    [...]
}
```