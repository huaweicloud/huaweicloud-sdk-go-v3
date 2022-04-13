package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckInstanceAccessResponse struct {
	// 返回值

	Result *bool `json:"result,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckInstanceAccessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckInstanceAccessResponse struct{}"
	}

	return strings.Join([]string{"CheckInstanceAccessResponse", string(data)}, " ")
}
