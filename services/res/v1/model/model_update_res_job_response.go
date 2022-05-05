package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateResJobResponse struct {

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回消息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateResJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateResJobResponse", string(data)}, " ")
}
