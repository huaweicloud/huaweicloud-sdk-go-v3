package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceQuery 查询条件参数类
type ResourceQuery struct {

	// 参数名
	Key string `json:"key"`

	// 参数值
	Value string `json:"value"`
}

func (o ResourceQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceQuery struct{}"
	}

	return strings.Join([]string{"ResourceQuery", string(data)}, " ")
}
