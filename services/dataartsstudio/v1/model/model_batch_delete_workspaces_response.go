package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteWorkspacesResponse Response Object
type BatchDeleteWorkspacesResponse struct {

	// 返回消息
	Message *string `json:"message,omitempty"`

	// 是否成功
	IsSuccess      *bool `json:"is_success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o BatchDeleteWorkspacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWorkspacesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteWorkspacesResponse", string(data)}, " ")
}
