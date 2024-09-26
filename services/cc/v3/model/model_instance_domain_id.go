package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceDomainId 网络实例（VPC，VGW）所属账号ID。
type InstanceDomainId struct {

	// 网络实例（VPC，VGW）所属账号ID。
	InstanceDomainId *string `json:"instance_domain_id,omitempty"`
}

func (o InstanceDomainId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDomainId struct{}"
	}

	return strings.Join([]string{"InstanceDomainId", string(data)}, " ")
}
