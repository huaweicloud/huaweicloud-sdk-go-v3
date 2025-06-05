package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyDto 委托结构体。
type AgencyDto struct {

	// 委托ID。
	Id *string `json:"id,omitempty"`

	// 委托名。
	Name *string `json:"name,omitempty"`

	// 被委托方账号名。
	TrustDomainName *string `json:"trust_domain_name,omitempty"`

	// 委托描述。
	Description *string `json:"description,omitempty"`

	// 委托权限列表。
	Roles *[]IamRoleDto `json:"roles,omitempty"`
}

func (o AgencyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyDto struct{}"
	}

	return strings.Join([]string{"AgencyDto", string(data)}, " ")
}
