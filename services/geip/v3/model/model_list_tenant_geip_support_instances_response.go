package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantGeipSupportInstancesResponse Response Object
type ListTenantGeipSupportInstancesResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 支持的Region对象
	SupportRegions *[]ListSupportInstancesSupportRegions `json:"support_regions,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTenantGeipSupportInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantGeipSupportInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListTenantGeipSupportInstancesResponse", string(data)}, " ")
}
