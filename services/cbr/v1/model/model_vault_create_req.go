package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultCreateReq struct {
	Vault *VaultCreate `json:"vault"`
}

func (o VaultCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultCreateReq struct{}"
	}

	return strings.Join([]string{"VaultCreateReq", string(data)}, " ")
}
