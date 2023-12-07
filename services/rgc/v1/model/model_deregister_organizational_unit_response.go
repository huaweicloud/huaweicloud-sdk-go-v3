package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeregisterOrganizationalUnitResponse Response Object
type DeregisterOrganizationalUnitResponse struct {

	// 异步接口的操作ID。
	OrganizationalUnitOperationId *string `json:"organizational_unit_operation_id,omitempty"`
	HttpStatusCode                int     `json:"-"`
}

func (o DeregisterOrganizationalUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeregisterOrganizationalUnitResponse struct{}"
	}

	return strings.Join([]string{"DeregisterOrganizationalUnitResponse", string(data)}, " ")
}
