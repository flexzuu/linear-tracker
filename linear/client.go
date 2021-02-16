package linear

import (
	"net/http"

	"github.com/hasura/go-graphql-client"
)

type authTransport struct {
	http.RoundTripper
	linearToken string
}

func (ct *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", ct.linearToken)
	return ct.RoundTripper.RoundTrip(req)
}

func NewClient(authToken string) *graphql.Client {
	httpClient := &http.Client{
		Transport: &authTransport{
			http.DefaultTransport,
			authToken,
		},
	}
	client := graphql.NewClient("https://api.linear.app/graphql", httpClient)
	return client
}
