package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigMaps struct {
	Configmap *ConfigMap `json:"configmap"`
}

func (o ConfigMaps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigMaps struct{}"
	}

	return strings.Join([]string{"ConfigMaps", string(data)}, " ")
}
