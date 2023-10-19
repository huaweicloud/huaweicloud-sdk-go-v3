package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LakeCatConfiguration 租户面配置项
type LakeCatConfiguration struct {

	// 配置项的key
	Key string `json:"key"`

	// 配置项的值
	Value string `json:"value"`

	// 配置项描述
	Description *string `json:"description,omitempty"`
}

func (o LakeCatConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LakeCatConfiguration struct{}"
	}

	return strings.Join([]string{"LakeCatConfiguration", string(data)}, " ")
}
