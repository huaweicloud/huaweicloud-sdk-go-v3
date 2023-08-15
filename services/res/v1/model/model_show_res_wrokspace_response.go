package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResWrokspaceResponse Response Object
type ShowResWrokspaceResponse struct {

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 工作空间详情。
	Workspaces *[]Workspaces `json:"workspaces,omitempty"`

	// 返回消息（请求成功时，不返回此字段）。
	Message *string `json:"message,omitempty"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResWrokspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResWrokspaceResponse struct{}"
	}

	return strings.Join([]string{"ShowResWrokspaceResponse", string(data)}, " ")
}
