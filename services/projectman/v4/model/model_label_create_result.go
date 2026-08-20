package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LabelCreateResult 工作项标签对象
type LabelCreateResult struct {

	// 标签ID，可通过查询标签列表接口获取，响应消息体中的id字段的值就是标签ID。 18~19个字符的数字字符串。
	Id *string `json:"id,omitempty"`

	// 标签所属工作项类型编码。
	CategoryTypes *[]string `json:"category_types,omitempty"`

	// 标签颜色RGB。 0~16个字符。
	Color *string `json:"color,omitempty"`

	// 标签标题。 2~256个字符。
	Title *string `json:"title,omitempty"`
}

func (o LabelCreateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelCreateResult struct{}"
	}

	return strings.Join([]string{"LabelCreateResult", string(data)}, " ")
}
