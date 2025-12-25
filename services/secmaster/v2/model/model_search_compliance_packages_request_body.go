package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCompliancePackagesRequestBody 查询基线检查遵从包列表请求体
type SearchCompliancePackagesRequestBody struct {

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 遵从包名称
	Name *string `json:"name,omitempty"`

	// 遵从包描述
	Description *string `json:"description,omitempty"`

	// 表示该遵从包的类型 0：内置 1: 自定义
	Type *int32 `json:"type,omitempty"`

	// 表示该遵从包的状态 0：停用 1: 启用
	State *int32 `json:"state,omitempty"`
}

func (o SearchCompliancePackagesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCompliancePackagesRequestBody struct{}"
	}

	return strings.Join([]string{"SearchCompliancePackagesRequestBody", string(data)}, " ")
}
