package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpgradeCandidateVersionsResponse Response Object
type ShowUpgradeCandidateVersionsResponse struct {

	// 升级类型信息列表
	UpgradeTypeList *[]UpgradeTypeInfo `json:"upgrade_type_list,omitempty"`

	// 是否可以回滚，true可以回滚，false不可以回滚。
	RollbackEnabled *bool `json:"rollback_enabled,omitempty"`

	// 原版本
	SourceVersion *string `json:"source_version,omitempty"`

	// 升级目标版本，没有在滚动升级中返回null。
	TargetVersion *string `json:"target_version,omitempty"`

	RollUpgradeProgress *RollUpgradeProgress `json:"roll_upgrade_progress,omitempty"`

	// 可以升级的版本，包括大小版本，滚动升级中返回空数组。
	UpgradeCandidateVersions *[]string `json:"upgrade_candidate_versions,omitempty"`

	// 可以升级的热补丁版本，滚动升级中返回空数组。
	HotfixUpgradeCandidateVersions *[]string `json:"hotfix_upgrade_candidate_versions,omitempty"`

	// 可以回滚的热补丁版本，滚动升级中返回空数组。
	HotfixRollbackCandidateVersions *[]string `json:"hotfix_rollback_candidate_versions,omitempty"`

	HotfixUpgradeInfos *HotfixUpgradeInfos `json:"hotfix_upgrade_infos,omitempty"`

	HotfixRollbackInfos *HotfixRollbackInfos `json:"hotfix_rollback_infos,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o ShowUpgradeCandidateVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpgradeCandidateVersionsResponse struct{}"
	}

	return strings.Join([]string{"ShowUpgradeCandidateVersionsResponse", string(data)}, " ")
}
