package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// flavor字段数据结构说明
type FlavorInfo struct {
	// 规格ID

	Id *string `json:"id,omitempty"`
	// 规格相关信息快捷链接

	Links *[]Links `json:"links,omitempty"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}
