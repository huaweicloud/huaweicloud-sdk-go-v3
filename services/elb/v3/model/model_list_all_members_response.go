package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllMembersResponse Response Object
type ListAllMembersResponse struct {

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// **参数解释**：后端服务器对象列表。
	Members        *[]MemberInfo `json:"members,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListAllMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllMembersResponse struct{}"
	}

	return strings.Join([]string{"ListAllMembersResponse", string(data)}, " ")
}
