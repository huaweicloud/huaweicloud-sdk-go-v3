package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMemberResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberResponse struct{}"
	}

	return strings.Join([]string{"ShowMemberResponse", string(data)}, " ")
}
