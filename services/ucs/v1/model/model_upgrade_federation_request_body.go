package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeFederationRequestBody struct {

	// 目标版本
	TargetVersion *string `json:"targetVersion,omitempty"`
}

func (o UpgradeFederationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeFederationRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeFederationRequestBody", string(data)}, " ")
}
