package model

import (
	"encoding/json"

	"strings"
)

type RegisterInfoRequestModel struct {
	// 大赛ID，大赛平台提供

	CompetitionId string `json:"competition_id"`
	// 大赛阶段ID，大赛平台提供

	StageId string `json:"stage_id"`
	// 租户ID

	DomainId string `json:"domain_id"`
}

func (o RegisterInfoRequestModel) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RegisterInfoRequestModel struct{}"
	}

	return strings.Join([]string{"RegisterInfoRequestModel", string(data)}, " ")
}
