package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionSampleInfo 原子动作样例详情。
type ActionSampleInfo struct {

	// 原子动作中文名称。
	ActionNameZh *string `json:"action_name_zh,omitempty"`

	// 原子动作英文名称。
	ActionNameEn *string `json:"action_name_en,omitempty"`

	// 动作Tag。
	ActionTag string `json:"action_tag"`

	// 动作分类名称。
	Catalog *string `json:"catalog,omitempty"`

	// 是否选择此动作。
	IsSelected *bool `json:"is_selected,omitempty"`

	// 原子动作样例文件下载地址。24小时内有效。
	SampleDownloadUrl *string `json:"sample_download_url,omitempty"`
}

func (o ActionSampleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionSampleInfo struct{}"
	}

	return strings.Join([]string{"ActionSampleInfo", string(data)}, " ")
}
