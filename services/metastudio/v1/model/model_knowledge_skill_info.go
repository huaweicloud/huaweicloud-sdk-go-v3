package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KnowledgeSkillInfo 知识库技能基本信息。
type KnowledgeSkillInfo struct {

	// 技能ID。
	SkillId *string `json:"skill_id,omitempty"`

	// 技能名称。
	Name *string `json:"name,omitempty"`

	// 技能标识。
	Identify *string `json:"identify,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o KnowledgeSkillInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KnowledgeSkillInfo struct{}"
	}

	return strings.Join([]string{"KnowledgeSkillInfo", string(data)}, " ")
}
