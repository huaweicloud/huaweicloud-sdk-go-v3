package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyGroupAndStreamResponse Response Object
type ModifyGroupAndStreamResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ModifyGroupAndStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyGroupAndStreamResponse struct{}"
	}

	return strings.Join([]string{"ModifyGroupAndStreamResponse", string(data)}, " ")
}
