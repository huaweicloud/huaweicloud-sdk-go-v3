package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTenantResourcesRequest Request Object
type CountTenantResourcesRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 业务类型。
	Business *string `json:"business,omitempty"`

	// 资源过期时间段 开始时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ResourceExpireStartTime *string `json:"resource_expire_start_time,omitempty"`

	// 资源过期时间段 结束时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ResourceExpireEndTime *string `json:"resource_expire_end_time,omitempty"`
}

func (o CountTenantResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTenantResourcesRequest struct{}"
	}

	return strings.Join([]string{"CountTenantResourcesRequest", string(data)}, " ")
}
