package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncParam struct {

	// 本次同步是否自动触发
	AutoTrigger bool `json:"auto_trigger"`

	// 混合云vault id
	VaultId string `json:"vault_id"`
}

func (o SyncParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncParam struct{}"
	}

	return strings.Join([]string{"SyncParam", string(data)}, " ")
}
