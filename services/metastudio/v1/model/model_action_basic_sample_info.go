package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionBasicSampleInfo 原子动作样例详情。
type ActionBasicSampleInfo struct {

	// 原子动作中文名称。
	ActionNameZh *string `json:"action_name_zh,omitempty"`

	// 原子动作英文名称。
	ActionNameEn *string `json:"action_name_en,omitempty"`

	// 原子动作标签。
	ActionTag string `json:"action_tag"`

	// 原子动作标签。
	Catalog *string `json:"catalog,omitempty"`

	// 是否选择此动作。
	IsSelected *bool `json:"is_selected,omitempty"`
}

func (o ActionBasicSampleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionBasicSampleInfo struct{}"
	}

	return strings.Join([]string{"ActionBasicSampleInfo", string(data)}, " ")
}
