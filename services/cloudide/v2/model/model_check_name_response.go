package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckNameResponse struct {
	// 返回值

	Result *bool `json:"result,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNameResponse struct{}"
	}

	return strings.Join([]string{"CheckNameResponse", string(data)}, " ")
}
