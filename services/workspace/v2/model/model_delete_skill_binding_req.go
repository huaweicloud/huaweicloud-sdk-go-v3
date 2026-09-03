package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSkillBindingReq 批量解绑技能请求。
type DeleteSkillBindingReq struct {

	// 实例 ID 列表。
	InstanceIds []string `json:"instance_ids"`

	// 技能 ID 列表。
	SkillIds []string `json:"skill_ids"`

	// 标签列表，格式为 key:value，通过标签查询 tbl_desktop_tags 表获取关联实例ID，与 instance_ids 合并后去重。
	Tags *[]string `json:"tags,omitempty"`
}

func (o DeleteSkillBindingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSkillBindingReq struct{}"
	}

	return strings.Join([]string{"DeleteSkillBindingReq", string(data)}, " ")
}
