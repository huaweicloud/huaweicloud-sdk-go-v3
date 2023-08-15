package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrganizationalUnitReqBody CreateOrganizationalUnit操作的请求体。
type CreateOrganizationalUnitReqBody struct {

	// 要分配给新组织单元的名称。
	Name string `json:"name"`

	// 要在其中创建新组织单元的根或组织单元的唯一标识符。
	ParentId string `json:"parent_id"`

	// 要附加到新创建的组织单元的标签列表。
	Tags *[]TagDto `json:"tags,omitempty"`
}

func (o CreateOrganizationalUnitReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationalUnitReqBody struct{}"
	}

	return strings.Join([]string{"CreateOrganizationalUnitReqBody", string(data)}, " ")
}
