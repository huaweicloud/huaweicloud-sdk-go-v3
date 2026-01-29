package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LabelEntity 工作项标签对象
type LabelEntity struct {

	// 标签id
	Id *string `json:"id,omitempty"`

	// 标签所属工作项类型，对应工作项的type字段
	LabelType *string `json:"label_type,omitempty"`

	// 标签颜色RGB
	Color *string `json:"color,omitempty"`

	// 标签标题
	Title *string `json:"title,omitempty"`
}

func (o LabelEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelEntity struct{}"
	}

	return strings.Join([]string{"LabelEntity", string(data)}, " ")
}
