package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAppQuotaBindableAppsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 客户端配额编号
	AppQuotaId string `json:"app_quota_id" xml:"app_quota_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 应用名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`
}

func (o ListAppQuotaBindableAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppQuotaBindableAppsRequest struct{}"
	}

	return strings.Join([]string{"ListAppQuotaBindableAppsRequest", string(data)}, " ")
}
