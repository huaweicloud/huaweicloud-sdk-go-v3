package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceDomainsResponse Response Object
type ListInstanceDomainsResponse struct {

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 域名信息
	Domains        *[]InstanceDomainItem `json:"domains,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListInstanceDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceDomainsResponse", string(data)}, " ")
}
