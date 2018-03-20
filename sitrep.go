package sitrep

import "github.com/sirupsen/logrus"

// This wrapper is being used to allow for predefined hooks
// used in all of my logging instances like Elasticsearch, Logstash, Syslog, File, Stdout

type Sitrep struct {
	*Logrus
}

func New() (*Sitrep, error) {
	return &Sitrep{}
}
