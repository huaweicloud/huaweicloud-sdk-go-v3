package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationStrategy 部署策略
type ConfigurationStrategy struct {

	// 部署策略
	Upgrade *string `json:"upgrade,omitempty"`
}

func (o ConfigurationStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationStrategy struct{}"
	}

	return strings.Join([]string{"ConfigurationStrategy", string(data)}, " ")
}
