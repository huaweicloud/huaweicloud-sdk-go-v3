package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ControlViolation 控制策略合规性。
type ControlViolation struct {

	// id。
	AccountId *string `json:"account_id,omitempty"`

	// name。
	AccountName *string `json:"account_name,omitempty"`

	// description。
	DisplayName *string `json:"display_name,omitempty"`

	// 控制策略名称。
	Name *string `json:"name,omitempty"`

	// 控制策略ID。
	ControlId *string `json:"control_id,omitempty"`

	// 父OU ID。
	ParentOrganizationUnitId *string `json:"parent_organization_unit_id,omitempty"`

	// 父OU名称。
	ParentOrganizationUnitName *string `json:"parent_organization_unit_name,omitempty"`

	// region。
	Region *string `json:"region,omitempty"`

	// 不合规资源。
	Resource *string `json:"resource,omitempty"`

	// 不合规资源类型。
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
