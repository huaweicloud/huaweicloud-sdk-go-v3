package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SkillPackageSummary 技能包摘要信息（用于技能详情中的 packages 列表）。
type SkillPackageSummary struct {

	// 技能包id。
	Id *string `json:"id,omitempty"`

	// 版本号。
	Version *string `json:"version,omitempty"`

	// 版本修订号。
	Revision *int32 `json:"revision,omitempty"`

	PackageStatus *PackageStatusEnum `json:"package_status,omitempty"`

	// 已部署的区域标识列表。
	Regions *[]string `json:"regions,omitempty"`
}

func (o SkillPackageSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SkillPackageSummary struct{}"
	}

	return strings.Join([]string{"SkillPackageSummary", string(data)}, " ")
}
