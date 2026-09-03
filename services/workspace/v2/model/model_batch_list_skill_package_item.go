package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListSkillPackageItem 批量查询技能包响应项。
type BatchListSkillPackageItem struct {

	// 技能 ID。
	SkillId *string `json:"skill_id,omitempty"`

	CurrentPackage *CurrentPackageInfo `json:"current_package,omitempty"`
}

func (o BatchListSkillPackageItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListSkillPackageItem struct{}"
	}

	return strings.Join([]string{"BatchListSkillPackageItem", string(data)}, " ")
}
