package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrganizationalUnitReqBody UpdateOrganizationalUnit 操作的请求体。
type UpdateOrganizationalUnitReqBody struct {

	// 要分配给新组织单元的名称。
	Name string `json:"name"`
}

func (o UpdateOrganizationalUnitReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrganizationalUnitReqBody struct{}"
	}

	return strings.Join([]string{"UpdateOrganizationalUnitReqBody", string(data)}, " ")
}
