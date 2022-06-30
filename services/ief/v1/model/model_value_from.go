package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValueFrom struct {
	Secret *Secrets `json:"secret,omitempty"`

	Configmap *ConfigsMap `json:"configmap,omitempty"`
}

func (o ValueFrom) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueFrom struct{}"
	}

	return strings.Join([]string{"ValueFrom", string(data)}, " ")
}
