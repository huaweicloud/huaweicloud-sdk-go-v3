package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommandsResponse Response Object
type ListCommandsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty"`

	// 命令列表
	Items          *[]Command `json:"items,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommandsResponse struct{}"
	}

	return strings.Join([]string{"ListCommandsResponse", string(data)}, " ")
}
