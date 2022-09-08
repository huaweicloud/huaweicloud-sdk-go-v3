package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页查询的公共属性
type Pages struct {

	// 页面起始页，从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 总数量。
	Count *int32 `json:"count,omitempty"`
}

func (o Pages) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pages struct{}"
	}

	return strings.Join([]string{"Pages", string(data)}, " ")
}
