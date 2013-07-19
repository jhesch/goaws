package goaws

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Route53 struct {
	auth *Auth
}

type ListHostedZonesResponse struct {
	XMLName     xml.Name `xml:"ListHostedZonesResponse"`
	HostedZones []HostedZones
	Marker      string
	IsTruncated bool
	NextMarker  string
	MaxItems    int
}

type GetHostedZoneResponse struct {
	XMLName       xml.Name `xml:"GetHostedZoneResponse"`
	HostedZone    []HostedZone
	DelegationSet DelegationSet
}

type HostedZones struct {
	XMLName    xml.Name `xml:"HostedZones"`
	HostedZone []HostedZone
}

type HostedZone struct {
	XMLName                xml.Name `xml:"HostedZone"`
	Id                     string
	Name                   string
	CallerReference        string
	Config                 Config
	ResourceRecordSetCount int
}

type Config struct {
	XMLName xml.Name `xml:"Config"`
	Comment string
}

type NameServers struct {
	NameServer []string
}

type DelegationSet struct {
	NameServers NameServers
}

var route53Endpoint = func() string {
	return "https://route53.amazonaws.com/2012-12-12/"
}()

// factory for the route53 type
func NewRoute53(AccessKey string, SecretKey string) *Route53 {
	auth := new(Auth)
	auth.setCredentials(AccessKey, SecretKey)
	r := new(Route53)
	r.auth = auth
	return r
}

// get a small subset of the hosted zones
func (r *Route53) getHostedZonesChunk(marker string, zones []HostedZone) []HostedZone {
	log.Printf("Get Hosted Zones Chunk: %s", marker)
	url := fmt.Sprintf("%shostedzone?maxitems=50", route53Endpoint)
	if marker != "" {
		url = fmt.Sprintf("%s&marker=%s", url, marker)
	}

	result, err := request(&RequestParams{Url: url, Auth: r.auth})

	v := ListHostedZonesResponse{}
	err = xml.Unmarshal(result, &v)
	if err != nil {
		return zones
	}
	for _, resultZones := range v.HostedZones {
		for _, zone := range resultZones.HostedZone {
			zones = append(zones, zone)
		}
	}

	if v.IsTruncated == true {
		zones = r.getHostedZonesChunk(v.NextMarker, zones)
	}
	return zones
}

// get a list of all hosted zones
func (r *Route53) GetHostedZones() []HostedZone {
	var hostedZones []HostedZone
	zones := r.getHostedZonesChunk("", hostedZones)

	return zones
}

// get a single hosted zone entity
func (r *Route53) GetHostedZone(id string) GetHostedZoneResponse {
	url := fmt.Sprintf("%s%s", route53Endpoint, id)

	result, err := request(&RequestParams{Url: url, Auth: r.auth})

	v := GetHostedZoneResponse{}
	err = xml.Unmarshal(result, &v)
	if err != nil {
		return v
	}

	return v
}
