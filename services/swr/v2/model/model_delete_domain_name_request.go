package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainNameRequest Request Object
type DeleteDomainNameRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 域名ID
	DomainnameId string `json:"domainname_id"`
}

func (o DeleteDomainNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainNameRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainNameRequest", string(data)}, " ")
}
