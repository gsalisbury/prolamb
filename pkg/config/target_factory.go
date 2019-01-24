package config

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	httpconfig "github.com/tlmiller/prolamb/pkg/http/config"
	"github.com/tlmiller/prolamb/pkg/target"
)

type TargetConfigFn func() interface{}

type TargetMapperFn func(name string, val interface{}) (target.Target, error)

type TargetMapping struct {
	Config TargetConfigFn
	Mapper TargetMapperFn
}

type TargetFactory map[string]TargetMapping

var (
	DefaultTargetFactory = make(TargetFactory)
)

func (t TargetFactory) AddMapping(id string, mapping TargetMapping) error {
	if _, exists := t[id]; exists {
		return fmt.Errorf("target mapping already registered for id '%s'", id)
	}
	t[id] = mapping
	return nil
}

func init() {
	httpMapping := TargetMapping{
		Config: httpconfig.New,
		Mapper: httpconfig.Mapper,
	}
	if err := DefaultTargetFactory.AddMapping("http", httpMapping); err != nil {
		panic(err)
	}
}

func (t TargetMapping) MakeTarget(name string, options json.RawMessage) (target.Target, error) {
	conf := t.Config()
	if err := json.Unmarshal(options, conf); err != nil {
		return nil, errors.Wrap(err, "making target mapping config from options")
	}
	return t.Mapper(name, conf)
}
