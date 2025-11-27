package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FederationVersionInfo struct {
	CurrentVersion *FederationVersionSpec `json:"currentVersion,omitempty"`

	// 目标联邦版本列表
	TargetVersions *[]FederationVersionSpec `json:"targetVersions,omitempty"`
}

func (o FederationVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FederationVersionInfo struct{}"
	}

	return strings.Join([]string{"FederationVersionInfo", string(data)}, " ")
}
