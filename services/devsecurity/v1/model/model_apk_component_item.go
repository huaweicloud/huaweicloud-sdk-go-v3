package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApkComponentItem struct {

	// 组件名称
	Name *string `json:"name,omitempty"`

	// 动作列表
	ActionList *[]string `json:"action_list,omitempty"`

	// 类别列表
	CategoryList *[]string `json:"category_list,omitempty"`

	// 主要活动
	IsMainActivity *bool `json:"is_main_activity,omitempty"`

	// 导出
	Exported *string `json:"exported,omitempty"`
}

func (o ApkComponentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApkComponentItem struct{}"
	}

	return strings.Join([]string{"ApkComponentItem", string(data)}, " ")
}
