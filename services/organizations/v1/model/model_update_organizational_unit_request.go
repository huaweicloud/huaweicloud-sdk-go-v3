package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrganizationalUnitRequest Request Object
type UpdateOrganizationalUnitRequest struct {

	// 与组织单元关联的唯一标识符（ID）。
	OrganizationalUnitId string `json:"organizational_unit_id"`

	Body *UpdateOrganizationalUnitReqBody `json:"body,omitempty"`
}

func (o UpdateOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"UpdateOrganizationalUnitRequest", string(data)}, " ")
}
