package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrganizationalUnitReqBody 创建注册OU的请求体。
type CreateOrganizationalUnitReqBody struct {

	// 要分配给新组织单元的名称。
	Name string `json:"name"`

	// 要在其中创建新组织单元的根或组织单元的唯一标识符。
	ParentId string `json:"parent_id"`
}

func (o CreateOrganizationalUnitReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationalUnitReqBody struct{}"
	}

	return strings.Join([]string{"CreateOrganizationalUnitReqBody", string(data)}, " ")
}
