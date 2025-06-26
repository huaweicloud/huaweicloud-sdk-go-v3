package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainNamesRequest Request Object
type ListDomainNamesRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 域名ID
	Uid *string `json:"uid,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`
}

func (o ListDomainNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainNamesRequest struct{}"
	}

	return strings.Join([]string{"ListDomainNamesRequest", string(data)}, " ")
}
