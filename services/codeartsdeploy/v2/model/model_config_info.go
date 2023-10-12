package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigInfo 部署任务执行参数
type ConfigInfo struct {

	// 参数名
	Name *string `json:"name,omitempty"`

	// 参数类型
	Type *string `json:"type,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`
}

func (o ConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigInfo struct{}"
	}

	return strings.Join([]string{"ConfigInfo", string(data)}, " ")
}
