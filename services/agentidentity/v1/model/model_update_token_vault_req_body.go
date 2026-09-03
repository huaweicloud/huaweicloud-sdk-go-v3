package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTokenVaultReqBody 请求体用于更新 TokenVault 配置。
type UpdateTokenVaultReqBody struct {
	PolicyEngineConfiguration *PolicyEngineConfiguration `json:"policy_engine_configuration"`
}

func (o UpdateTokenVaultReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTokenVaultReqBody struct{}"
	}

	return strings.Join([]string{"UpdateTokenVaultReqBody", string(data)}, " ")
}
