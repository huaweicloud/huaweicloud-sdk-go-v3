package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptsRequest Request Object
type ListScriptsRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 脚本ID
	Id *string `json:"id,omitempty"`

	// 脚本名称。
	Name *string `json:"name,omitempty"`

	// 脚本类型。
	Type *string `json:"type,omitempty"`
}

func (o ListScriptsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptsRequest", string(data)}, " ")
}
