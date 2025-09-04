package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigInfo 配置信息
type ConfigInfo struct {
	Url *TopUrl `json:"url,omitempty"`

	Ua *TopUa `json:"ua,omitempty"`
}

func (o ConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigInfo struct{}"
	}

	return strings.Join([]string{"ConfigInfo", string(data)}, " ")
}
