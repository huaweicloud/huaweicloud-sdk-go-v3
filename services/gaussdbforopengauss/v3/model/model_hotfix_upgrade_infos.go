package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HotfixUpgradeInfos 可以升级的热补丁信息。
type HotfixUpgradeInfos struct {

	// 热补丁版本。
	Version *string `json:"version,omitempty"`

	// 通用/非通用补丁信息。 枚举值：   \"common\": 通用补丁。   \"certain\": 定制补丁。
	CommonPatch *string `json:"common_patch,omitempty"`

	// 是否和备份相关。
	BackupSensitive *bool `json:"backup_sensitive,omitempty"`

	// 补丁的描述信息。
	Descripition *string `json:"descripition,omitempty"`
}

func (o HotfixUpgradeInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotfixUpgradeInfos struct{}"
	}

	return strings.Join([]string{"HotfixUpgradeInfos", string(data)}, " ")
}
