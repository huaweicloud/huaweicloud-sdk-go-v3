package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OnPremisesConfig struct {

	// 本地集群版本信息列表
	VersionsInfo *[]OnPremisesVersionInfo `json:"versionsInfo,omitempty"`
}

func (o OnPremisesConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OnPremisesConfig struct{}"
	}

	return strings.Join([]string{"OnPremisesConfig", string(data)}, " ")
}
