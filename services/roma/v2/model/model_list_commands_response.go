package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCommandsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 命令列表
	Items          *[]Command `json:"items,omitempty" xml:"items"`
	HttpStatusCode int        `json:"-"`
}

func (o ListCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommandsResponse struct{}"
	}

	return strings.Join([]string{"ListCommandsResponse", string(data)}, " ")
}
