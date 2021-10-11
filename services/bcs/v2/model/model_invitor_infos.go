package model

import (
	"encoding/json"

	"strings"
)

// 邀请方信息
type InvitorInfos struct {
	// 邀请方租户ID

	TenantId string `json:"tenant_id"`
	// 邀请方项目ID

	ProjectId string `json:"project_id"`
	// 邀请方服务实例ID

	BlockchainId string `json:"blockchain_id"`
}

func (o InvitorInfos) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InvitorInfos struct{}"
	}

	return strings.Join([]string{"InvitorInfos", string(data)}, " ")
}
