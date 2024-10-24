package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantResourcesRequest Request Object
type ListTenantResourcesRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 资源类型。
	Resource *string `json:"resource,omitempty"`

	// 业务类型。
	Business *string `json:"business,omitempty"`

	// 资源来源。
	ResourceSource string `json:"resource_source"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源id。
	ResourceId *string `json:"resource_id,omitempty"`

	// 订单id。
	OrderId *string `json:"order_id,omitempty"`

	// 计费模式。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 资源过期时间段 开始时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ResourceExpireStartTime *string `json:"resource_expire_start_time,omitempty"`

	// 资源过期时间段 结束时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ResourceExpireEndTime *string `json:"resource_expire_end_time,omitempty"`

	// 子资源类型。
	SubResource *string `json:"sub_resource,omitempty"`
}

func (o ListTenantResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListTenantResourcesRequest", string(data)}, " ")
}
