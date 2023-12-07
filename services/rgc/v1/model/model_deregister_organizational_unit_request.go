package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeregisterOrganizationalUnitRequest Request Object
type DeregisterOrganizationalUnitRequest struct {

	// OU ID。
	ManagedOrganizationUnitId string `json:"managed_organization_unit_id"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o DeregisterOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeregisterOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"DeregisterOrganizationalUnitRequest", string(data)}, " ")
}
