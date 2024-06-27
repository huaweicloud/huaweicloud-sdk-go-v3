package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HotfixVersionInfo 热补丁信息。
type HotfixVersionInfo struct {

	// 热补丁版本。
	Version string `json:"version"`

	// 热补丁升级完成时间列表。  热补丁升级完成时间，格式为“yyyy-mm-dd hh:mm:ss timezone”。  其中timezone是指时区。
	UpgradeFinishedTime *string `json:"upgrade_finished_time,omitempty"`
}

func (o HotfixVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotfixVersionInfo struct{}"
	}

	return strings.Join([]string{"HotfixVersionInfo", string(data)}, " ")
}
