package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValueFrom struct {
	Configmap *ConfigsMap `json:"configmap,omitempty"`

	Secret *DeploymentSecrets `json:"secret,omitempty"`
}

func (o ValueFrom) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueFrom struct{}"
	}

	return strings.Join([]string{"ValueFrom", string(data)}, " ")
}
