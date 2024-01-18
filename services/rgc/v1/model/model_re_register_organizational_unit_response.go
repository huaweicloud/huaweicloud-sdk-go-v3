package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReRegisterOrganizationalUnitResponse Response Object
type ReRegisterOrganizationalUnitResponse struct {

	// 异步接口的操作ID。
	OrganizationalUnitOperationId *string `json:"organizational_unit_operation_id,omitempty"`
	HttpStatusCode                int     `json:"-"`
}

func (o ReRegisterOrganizationalUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReRegisterOrganizationalUnitResponse struct{}"
	}

	return strings.Join([]string{"ReRegisterOrganizationalUnitResponse", string(data)}, " ")
}
