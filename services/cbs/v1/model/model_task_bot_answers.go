package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type TaskBotAnswers struct {

	// 答案。
	Answer string `json:"answer"`

	// 技能ID。
	SkillId string `json:"skill_id"`

	// 技能回复信息。
	SkillResponses []SkillResponse `json:"skill_responses"`
}

func (o TaskBotAnswers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBotAnswers struct{}"
	}

	return strings.Join([]string{"TaskBotAnswers", string(data)}, " ")
}
