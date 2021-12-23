package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 存储库修改参数体
type VaultUpdateReq struct {
	Vault *VaultUpdate `json:"vault"`
}

func (o VaultUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultUpdateReq struct{}"
	}

	return strings.Join([]string{"VaultUpdateReq", string(data)}, " ")
}
