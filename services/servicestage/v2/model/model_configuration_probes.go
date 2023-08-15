package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationProbes struct {
	LivenessProbe *ComponentProbe `json:"livenessProbe,omitempty"`

	ReadinessProbe *ComponentProbe `json:"readinessProbe,omitempty"`
}

func (o ConfigurationProbes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationProbes struct{}"
	}

	return strings.Join([]string{"ConfigurationProbes", string(data)}, " ")
}
