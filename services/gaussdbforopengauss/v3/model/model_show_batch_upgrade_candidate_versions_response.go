package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchUpgradeCandidateVersionsResponse Response Object
type ShowBatchUpgradeCandidateVersionsResponse struct {

	// 升级类型信息列表。
	UpgradeTypeList *[]UpgradeTypeInfo `json:"upgrade_type_list,omitempty"`

	// 升级目标版本，没有在滚动升级中返回null。
	TargetVersion *string `json:"target_version,omitempty"`

	// 可以升级的版本，包括大小版本。
	UpgradeCandidateVersions *[]string `json:"upgrade_candidate_versions,omitempty"`

	// 可以升级的热补丁信息。
	HotfixUpgradeInfos *[]HotfixInfo `json:"hotfix_upgrade_infos,omitempty"`

	// 可以回滚的热补丁信息。
	HotfixRollbackInfos *[]HotfixInfo `json:"hotfix_rollback_infos,omitempty"`
	HttpStatusCode      int           `json:"-"`
}

func (o ShowBatchUpgradeCandidateVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchUpgradeCandidateVersionsResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchUpgradeCandidateVersionsResponse", string(data)}, " ")
}
