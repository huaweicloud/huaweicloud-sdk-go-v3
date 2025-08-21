package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationTreeResponse Response Object
type ListOrganizationTreeResponse struct {

	// **参数解释**： 查询组织账号响应 **取值范围**： 不涉及
	Data *[]OrganizationAccountInfo `json:"data,omitempty"`

	// **参数解释**： 分页标记 **取值范围**： 不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOrganizationTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationTreeResponse struct{}"
	}

	return strings.Join([]string{"ListOrganizationTreeResponse", string(data)}, " ")
}
