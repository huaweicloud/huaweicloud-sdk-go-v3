package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationInfoDto struct {

	// 节点配置信息
	Configuration *[]ConfigurationDetail `json:"configuration,omitempty"`
}

func (o ConfigurationInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationInfoDto struct{}"
	}

	return strings.Join([]string{"ConfigurationInfoDto", string(data)}, " ")
}
