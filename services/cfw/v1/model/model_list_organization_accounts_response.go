package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationAccountsResponse Response Object
type ListOrganizationAccountsResponse struct {

	// **参数解释**： 查询组织账号响应 **取值范围**： 不涉及
	Data *[]OrganizationAccountInfo `json:"data,omitempty"`

	// **参数解释**： 分页标记 **取值范围**： 不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOrganizationAccountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationAccountsResponse struct{}"
	}

	return strings.Join([]string{"ListOrganizationAccountsResponse", string(data)}, " ")
}
