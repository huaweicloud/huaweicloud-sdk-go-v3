package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnrollAccountRequestBody 纳管账号。
type EnrollAccountRequestBody struct {

	// 注册OU 标识。
	ParentOrganizationUnitId string `json:"parent_organization_unit_id"`

	Blueprint *Blueprint `json:"blueprint,omitempty"`
}

func (o EnrollAccountRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnrollAccountRequestBody struct{}"
	}

	return strings.Join([]string{"EnrollAccountRequestBody", string(data)}, " ")
}
