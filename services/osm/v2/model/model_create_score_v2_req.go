package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateScoreV2Req struct {
	// 评价内容

	Judgement string `json:"judgement"`
	// 满意度列表

	IncidentSatisfactionList []IncidentSatisfactionV2Do `json:"incident_satisfaction_list"`
	// 组id

	GroupId *string `json:"group_id,omitempty"`
}

func (o CreateScoreV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScoreV2Req struct{}"
	}

	return strings.Join([]string{"CreateScoreV2Req", string(data)}, " ")
}
