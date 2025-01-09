package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantConfigsResponse Response Object
type ListTenantConfigsResponse struct {

	// 租户个性配置列表
	FunctionConfigs *[]FunctionConfig `json:"function_configs,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o ListTenantConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListTenantConfigsResponse", string(data)}, " ")
}
