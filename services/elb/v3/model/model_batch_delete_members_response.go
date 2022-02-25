package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteMembersResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`
	// 后端服务器对象列表。

	Members        *[]BatchDeleteMemberState `json:"members,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o BatchDeleteMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersResponse", string(data)}, " ")
}
