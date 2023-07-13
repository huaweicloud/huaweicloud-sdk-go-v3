package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LabelPageListDto 标签页面列表返回的单个标签对象
type LabelPageListDto struct {

	// 标签页面id
	Id *string `json:"id,omitempty"`

	// 标签页面标题
	Name *string `json:"name,omitempty"`

	// 标签页面类型
	Feature *string `json:"feature,omitempty"`

	// 标签页面包含的标签
	Labels *[]string `json:"labels,omitempty"`

	// 标签页面创建者
	Creator *string `json:"creator,omitempty"`
}

func (o LabelPageListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelPageListDto struct{}"
	}

	return strings.Join([]string{"LabelPageListDto", string(data)}, " ")
}
