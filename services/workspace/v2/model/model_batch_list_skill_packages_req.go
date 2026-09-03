package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListSkillPackagesReq 批量查询技能包请求。
type BatchListSkillPackagesReq struct {

	// 技能 ID 列表。
	SkillIds []string `json:"skill_ids"`
}

func (o BatchListSkillPackagesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListSkillPackagesReq struct{}"
	}

	return strings.Join([]string{"BatchListSkillPackagesReq", string(data)}, " ")
}
