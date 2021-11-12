package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ConsumeDeadlettersMessageResponse struct {
	// 消息数组。

	Body           *[]ConsumeDeadlettersMessage `json:"body,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ConsumeDeadlettersMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumeDeadlettersMessageResponse struct{}"
	}

	return strings.Join([]string{"ConsumeDeadlettersMessageResponse", string(data)}, " ")
}
