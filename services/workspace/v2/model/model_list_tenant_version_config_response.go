package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantVersionConfigResponse Response Object
type ListTenantVersionConfigResponse struct {

	// 总记录数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 版本列表
	Records        *[]TenantVersionInfo `json:"records,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListTenantVersionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantVersionConfigResponse struct{}"
	}

	return strings.Join([]string{"ListTenantVersionConfigResponse", string(data)}, " ")
}
