package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FederationConfig struct {

	// 联邦版本信息列表
	VersionsInfo *[]FederationVersionInfo `json:"versionsInfo,omitempty"`
}

func (o FederationConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FederationConfig struct{}"
	}

	return strings.Join([]string{"FederationConfig", string(data)}, " ")
}
