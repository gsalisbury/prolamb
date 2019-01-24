package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/tlmiller/prolamb/pkg/target"
)

type TargetConfig struct {
	Name    string          `json:"name"`
	Type    string          `json:"type"`
	Options json.RawMessage `json:"options"`
}

func GetTargetsFromRawQuery(val string) (target.Target, error) {
	conf := TargetConfig{}
	jsonTarget, err := base64.URLEncoding.DecodeString(val)
	if err != nil {
		return nil, errors.Wrap(err, "decoding target base64 query")
	}
	if err := json.Unmarshal(jsonTarget, &conf); err != nil {
		return nil, errors.Wrap(err, "decoding target json from query")
	}

	if conf.Name == "" {
		return nil, errors.New("target name cannot be empty")
	}
	if conf.Type == "" {
		return nil, errors.New("target type cannot be empty")
	}

	mapping, exists := DefaultTargetFactory[conf.Type]
	if !exists {
		return nil, fmt.Errorf("no target mapping for %s", conf.Type)
	}
	target, err := mapping.MakeTarget(conf.Name, conf.Options)
	if err != nil {
		return nil, errors.Wrapf(err, "making target for type %s and name %s",
			conf.Type, conf.Name)
	}
	return target, nil
}
