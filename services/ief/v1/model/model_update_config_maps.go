package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConfigMaps struct {
	Configmap *UpdateConfigMap `json:"configmap"`
}

func (o UpdateConfigMaps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigMaps struct{}"
	}

	return strings.Join([]string{"UpdateConfigMaps", string(data)}, " ")
}
