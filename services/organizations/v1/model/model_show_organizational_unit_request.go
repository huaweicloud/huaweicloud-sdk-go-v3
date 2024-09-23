package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationalUnitRequest Request Object
type ShowOrganizationalUnitRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 与组织单元关联的唯一标识符（ID）。
	OrganizationalUnitId string `json:"organizational_unit_id"`
}

func (o ShowOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"ShowOrganizationalUnitRequest", string(data)}, " ")
}
