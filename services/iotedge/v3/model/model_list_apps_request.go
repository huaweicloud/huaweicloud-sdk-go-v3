package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAppsRequest struct {

	// 应用类型
	AppType *string `json:"app_type,omitempty"`

	// 应用来源
	ProviderType *string `json:"provider_type,omitempty"`

	// 应用ID
	AppId *string `json:"app_id,omitempty"`

	// 每页记录数，默认值为10，取值区间为1-1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsRequest struct{}"
	}

	return strings.Join([]string{"ListAppsRequest", string(data)}, " ")
}
