package model

import (
	"encoding/json"

	"strings"
)

// 存储库修改参数体
type VaultUpdateReq struct {
	Vault *VaultUpdate `json:"vault"`
}

func (o VaultUpdateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VaultUpdateReq struct{}"
	}

	return strings.Join([]string{"VaultUpdateReq", string(data)}, " ")
}
