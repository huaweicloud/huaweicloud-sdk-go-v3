package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionTagInfo 分身动作Tag映射信息。
type ActionTagInfo struct {

	// 原子动作中文名称。
	ActionNameZh string `json:"action_name_zh"`

	// 原子动作英文名称。
	ActionNameEn string `json:"action_name_en"`

	// 动作分类名称。
	Catalog *string `json:"catalog,omitempty"`

	// 样例视频文件名，最大长度256，最小长度1。
	FileName *string `json:"file_name,omitempty"`

	// 动作标签。
	Tag *string `json:"tag,omitempty"`
}

func (o ActionTagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionTagInfo struct{}"
	}

	return strings.Join([]string{"ActionTagInfo", string(data)}, " ")
}
