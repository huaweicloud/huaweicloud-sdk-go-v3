package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LatestVersionInfo HDP最新版本信息。
type LatestVersionInfo struct {

	// 租户的HDP最新版本。
	LatestVersion *string `json:"latest_version,omitempty"`

	// HDA类型： * `SBC` - 非VDI下SBC类型 * `OA_APP` - VDI下非GPU类型 * `PT_APP` - VDI下GPU类型
	HdaType *string `json:"hda_type,omitempty"`
}

func (o LatestVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LatestVersionInfo struct{}"
	}

	return strings.Join([]string{"LatestVersionInfo", string(data)}, " ")
}
