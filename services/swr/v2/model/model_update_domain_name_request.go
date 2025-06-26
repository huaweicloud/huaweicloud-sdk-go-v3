package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainNameRequest Request Object
type UpdateDomainNameRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 域名ID
	DomainnameId string `json:"domainname_id"`

	Body *UpdateDomainNameRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainNameRequest", string(data)}, " ")
}
