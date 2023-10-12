package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTagParam 标签查询参数
type ResourceTagParam struct {

	// 资源标签的键
	Key string `json:"key"`

	// 资源标签的值列表
	Values *[]string `json:"values,omitempty"`
}

func (o ResourceTagParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagParam struct{}"
	}

	return strings.Join([]string{"ResourceTagParam", string(data)}, " ")
}
