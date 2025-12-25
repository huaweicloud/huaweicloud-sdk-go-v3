package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCheckitemsResponse Response Object
type SearchCheckitemsResponse struct {

	// 内置检查项个数
	BuiltinCheckitemNum float32 `json:"builtin_checkitem_num,omitempty"`

	// 检查项总数
	CheckitemNum float32 `json:"checkitem_num,omitempty"`

	// 自定义检查项个数
	CustomizedCheckitemNum float32 `json:"customized_checkitem_num,omitempty"`

	// 检查项总数
	Count float32 `json:"count,omitempty"`

	// 检查项详情
	Checkitems     *[]CheckitemModel `json:"checkitems,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o SearchCheckitemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCheckitemsResponse struct{}"
	}

	return strings.Join([]string{"SearchCheckitemsResponse", string(data)}, " ")
}
