package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValueFrom struct {
	Secret *Secrets `json:"secret,omitempty" xml:"secret"`

	Configmap *ConfigsMap `json:"configmap,omitempty" xml:"configmap"`
}

func (o ValueFrom) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueFrom struct{}"
	}

	return strings.Join([]string{"ValueFrom", string(data)}, " ")
}
