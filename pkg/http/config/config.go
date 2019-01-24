package config

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	httptarget "github.com/tlmiller/prolamb/pkg/http"
	"github.com/tlmiller/prolamb/pkg/target"
)

type Config struct {
	Method string
	URL    string
}

func New() interface{} {
	return &Config{
		Method: http.MethodGet,
	}
}

func Mapper(name string, val interface{}) (target.Target, error) {
	conf, ok := val.(*Config)
	if !ok {
		return nil, errors.New("unknown type for mapper config")
	}

	if conf.URL == "" {
		return nil, errors.New("http target url cannot be empty")
	}
	if conf.Method == "" {
		return nil, errors.New("http target method cannot be empty")
	}

	url, err := url.Parse(conf.URL)
	if err != nil {
		return nil, errors.Wrap(err, "http target parsing url")
	}

	return &httptarget.Target{
		Name:   name,
		Method: conf.Method,
		URL:    url,
	}, nil
}
