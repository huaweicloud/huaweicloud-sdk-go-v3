package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ControlViolation 控制策略合规性。
type ControlViolation struct {

	// 控制策略纳管账号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 控制策略纳管账号名称。
	AccountName *string `json:"account_name,omitempty"`

	// 控制策略显示名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 控制策略名称。
	Name *string `json:"name,omitempty"`

	// 控制策略ID。
	ControlId *string `json:"control_id,omitempty"`

	// 父注册OU ID。
	ParentOrganizationalUnitId *string `json:"parent_organizational_unit_id,omitempty"`

	// 父注册OU名称。
	ParentOrganizationalUnitName *string `json:"parent_organizational_unit_name,omitempty"`

	// 区域名称。
	Region *string `json:"region,omitempty"`

	// 控制策略不合规资源。
	Resource *string `json:"resource,omitempty"`

	// 控制策略不合规资源的名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 控制策略不合规资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 云服务名称。
	Service *string `json:"service,omitempty"`
}

func (o ControlViolation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ControlViolation struct{}"
	}

	return strings.Join([]string{"ControlViolation", string(data)}, " ")
}
