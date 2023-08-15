package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMemberResponse Response Object
type UpdateMemberResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberResponse struct{}"
	}

	return strings.Join([]string{"UpdateMemberResponse", string(data)}, " ")
}
