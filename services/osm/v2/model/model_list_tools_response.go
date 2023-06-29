package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListToolsResponse Response Object
type ListToolsResponse struct {

	// 推荐工具列表
	Tools *[]Tool `json:"tools,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListToolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListToolsResponse struct{}"
	}

	return strings.Join([]string{"ListToolsResponse", string(data)}, " ")
}
