package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RdsUpgradePrecheckV3Req mysql5.7升级mysql8.0预检查请求体
type RdsUpgradePrecheckV3Req struct {

	// 特定场景使用，普通用户无需填写该字段。
	TargetVersion *string `json:"target_version,omitempty"`
}

func (o RdsUpgradePrecheckV3Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdsUpgradePrecheckV3Req struct{}"
	}

	return strings.Join([]string{"RdsUpgradePrecheckV3Req", string(data)}, " ")
}
