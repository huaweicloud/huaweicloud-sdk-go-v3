package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateMembersResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`
	// 后端服务器对象列表。

	Members        *[]BatchMember `json:"members,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o BatchCreateMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersResponse", string(data)}, " ")
}
