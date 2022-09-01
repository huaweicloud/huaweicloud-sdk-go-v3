package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVpcsRequest struct {

	// 查询返回虚拟私有云列表数量。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 通过ID查询
	Id *string `json:"id,omitempty" xml:"id"`

	// 通过name查询
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o ListVpcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcsRequest", string(data)}, " ")
}
