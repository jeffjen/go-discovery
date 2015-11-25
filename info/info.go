package info

import (
	disc "github.com/jeffjen/go-discovery"

	_ "github.com/Sirupsen/logrus"

	"encoding/json"
	"net/http"
	"os"
	"time"
)

var (
	VERSION = os.Getenv("VERSION")

	BUILD = os.Getenv("BUILD")

	NODE_NAME = os.Getenv("NODE_NAME")

	NODE_REGION = os.Getenv("NODE_REGION")

	NODE_AVAIL_ZONE = os.Getenv("NODE_AVAIL_ZONE")

	NODE_PUBLIC_HOSTNAME = os.Getenv("NODE_PUBLIC_HOSTNAME")

	NODE_PUBLIC_IPV4 = os.Getenv("NODE_PUBLIC_IPV4")

	NODE_PRIVATE_IPV4 = os.Getenv("NODE_PRIVATE_IPV4")

	MetaData string
)

type NodeInfo struct {
	Version   string `json:"version"`
	Build     string `json:"build"`
	Node      string `json:"node,omitempty"`
	Region    string `json:"region,omitempty"`
	Zone      string `json:"avail_zone,omitempty"`
	Host      string `json:"host,omitempty"`
	Public    string `json:"public_ipv4,omitempty"`
	Private   string `json:"local_ipv4,omitempty"`
	Discovery string `json:"discovery"`
	Hearbeat  string `json:"heartbeat"`
	TTL       string `json:"ttl"`
	Timestamp string `json:"current_time"`
}

func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(NodeInfo{
		Version:   VERSION,
		Build:     BUILD,
		Node:      NODE_NAME,
		Region:    NODE_REGION,
		Zone:      NODE_AVAIL_ZONE,
		Host:      NODE_PUBLIC_HOSTNAME,
		Public:    NODE_PUBLIC_IPV4,
		Private:   NODE_PRIVATE_IPV4,
		Discovery: disc.Discovery,
		Hearbeat:  disc.Hearbeat.String(),
		TTL:       disc.TTL.String(),
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func init() {
	type metadata struct {
		Node    string `json:"node,omitempty"`
		Region  string `json:"region,omitempty"`
		Zone    string `json:"avail_zone,omitempty"`
		Host    string `json:"host,omitempty"`
		Public  string `json:"public_ipv4,omitempty"`
		Private string `json:"local_ipv4,omitempty"`
	}
	b, _ := json.Marshal(metadata{
		NODE_NAME,
		NODE_REGION,
		NODE_AVAIL_ZONE,
		NODE_PUBLIC_HOSTNAME,
		NODE_PUBLIC_IPV4,
		NODE_PRIVATE_IPV4,
	})
	MetaData = string(b)
}
