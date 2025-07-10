package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateManagedAccountRequestBody 更新账号的基本信息。
type UpdateManagedAccountRequestBody struct {

	// 父注册OU ID。
	ParentOrganizationalUnitId string `json:"parent_organizational_unit_id"`

	// 父注册OU名称。
	ParentOrganizationalUnitName string `json:"parent_organizational_unit_name"`

	Blueprint *Blueprint `json:"blueprint,omitempty"`
}

func (o UpdateManagedAccountRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateManagedAccountRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateManagedAccountRequestBody", string(data)}, " ")
}
