package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDdmConfigurationsResponse Response Object
type ListDdmConfigurationsResponse struct {

	// 参数配置列表
	Configurations *[]ConfigurationInfo `json:"configurations,omitempty"`

	// 参数模板总数。
	Total *int32 `json:"total,omitempty"`

	// 分页参数: 起始值。
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数：每页多少条。
	Limit          *int32 `json:"limit,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDdmConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDdmConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListDdmConfigurationsResponse", string(data)}, " ")
}
