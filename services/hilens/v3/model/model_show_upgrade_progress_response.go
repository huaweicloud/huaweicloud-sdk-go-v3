package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpgradeProgressResponse Response Object
type ShowUpgradeProgressResponse struct {

	// 升级的固件版本
	Version *string `json:"version,omitempty"`

	// 固件升级状态，1:升级中 2:升级失败 3:升级成功
	Status *int32 `json:"status,omitempty"`

	// 固件升级进度，用数字0-100表示
	Progress *string `json:"progress,omitempty"`

	// 固件升级失败原因
	Cause          *string `json:"cause,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUpgradeProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpgradeProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowUpgradeProgressResponse", string(data)}, " ")
}
