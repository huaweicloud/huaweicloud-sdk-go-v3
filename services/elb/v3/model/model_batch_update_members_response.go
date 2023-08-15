package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateMembersResponse Response Object
type BatchUpdateMembersResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	// 后端服务器对象列表。
	Members        *[]BatchUpdateMember `json:"members,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o BatchUpdateMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersResponse", string(data)}, " ")
}
