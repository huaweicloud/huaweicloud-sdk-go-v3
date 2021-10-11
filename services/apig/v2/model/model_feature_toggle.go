package model

import (
	"encoding/json"

	"strings"
)

type FeatureToggle struct {
	// 特性名称

	Name string `json:"name"`
	// 是否开启特性

	Enable bool `json:"enable"`
	// 特性参数配置

	Config *string `json:"config,omitempty"`
}

func (o FeatureToggle) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FeatureToggle struct{}"
	}

	return strings.Join([]string{"FeatureToggle", string(data)}, " ")
}
