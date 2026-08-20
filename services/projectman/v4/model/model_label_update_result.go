package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LabelUpdateResult 工作项标签对象
type LabelUpdateResult struct {

	// 标签所属工作项类型编码。
	CategoryTypes *[]string `json:"category_types,omitempty"`

	// 标签颜色RGB。 0~16个字符。
	Color *string `json:"color,omitempty"`

	// 标签标题。 2~256个字符。
	Title *string `json:"title,omitempty"`
}

func (o LabelUpdateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelUpdateResult struct{}"
	}

	return strings.Join([]string{"LabelUpdateResult", string(data)}, " ")
}
