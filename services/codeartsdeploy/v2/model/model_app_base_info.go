package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppBaseInfo 应用基本信息
type AppBaseInfo struct {

	// 应用id
	Id *string `json:"id,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 局点信息
	Region *string `json:"region,omitempty"`

	// 应用是否被禁用
	IsDisable *bool `json:"is_disable,omitempty"`
}

func (o AppBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppBaseInfo struct{}"
	}

	return strings.Join([]string{"AppBaseInfo", string(data)}, " ")
}
