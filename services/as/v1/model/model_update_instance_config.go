package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceConfig struct {

	// 镜像ID
	ImageRef string `json:"imageRef"`
}

func (o UpdateInstanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfig struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfig", string(data)}, " ")
}
