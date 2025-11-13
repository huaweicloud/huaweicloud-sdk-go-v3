package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInsertCommandsResponse Response Object
type ListInsertCommandsResponse struct {

	// 命令总数。
	Count *int32 `json:"count,omitempty"`

	// 插入播放命令列表。
	InsertCommands *[]InsertCommandItem `json:"insert_commands,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInsertCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInsertCommandsResponse struct{}"
	}

	return strings.Join([]string{"ListInsertCommandsResponse", string(data)}, " ")
}
