package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSkillBindingReq 批量绑定技能到实例请求。
type CreateSkillBindingReq struct {

	// 实例 ID 列表。
	InstanceIds []string `json:"instance_ids"`

	// 技能 ID 列表。
	SkillIds []string `json:"skill_ids"`

	// 技能版本列表，与 skill_ids 一一对应。不传或对应位置为空时使用技能当前版本。
	Versions *[]string `json:"versions,omitempty"`

	// 标签列表，格式为 key:value，通过标签查询 tbl_desktop_tags 表获取关联实例ID，与 instance_ids 合并后去重。
	Tags *[]string `json:"tags,omitempty"`
}

func (o CreateSkillBindingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSkillBindingReq struct{}"
	}

	return strings.Join([]string{"CreateSkillBindingReq", string(data)}, " ")
}
