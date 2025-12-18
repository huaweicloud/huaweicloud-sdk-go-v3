package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgencyInfosResponse Response Object
type ListAgencyInfosResponse struct {

	// **参数解释**: 委托名称。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 委托是否存在。 **取值范围**: true、false
	IsExisted *bool `json:"is_existed,omitempty"`

	// **参数解释**: 委托绑定的权限策略信息。
	Roles          *[]AgencyRoleResult `json:"roles,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListAgencyInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgencyInfosResponse struct{}"
	}

	return strings.Join([]string{"ListAgencyInfosResponse", string(data)}, " ")
}
