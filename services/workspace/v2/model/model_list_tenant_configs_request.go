package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantConfigsRequest Request Object
type ListTenantConfigsRequest struct {

	// 开关名称。
	Name *string `json:"name,omitempty"`
}

func (o ListTenantConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListTenantConfigsRequest", string(data)}, " ")
}
