package goaws

import (
  "fmt"
)

type Route53 struct {
  auth *Auth
}

var route53Endpoint = func() string {
  return "https://route53.amazonaws.com/2012-12-12/"
}()

func NewRoute53(AccessKey string, SecretKey string) *Route53 {
	auth := new(Auth)
	auth.setCredentials(AccessKey, SecretKey)
  r := new(Route53)
  r.auth = auth
  return r
}

func (r *Route53) ListHostedZones() {
  url := fmt.Sprintf("%shostedzone", route53Endpoint)
  result, err := Request(&RequestParams{Url: url, Auth: r.auth})
  fmt.Printf("%+v - %+v", result, err)
}
