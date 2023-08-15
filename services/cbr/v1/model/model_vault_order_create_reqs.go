package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VaultOrderCreateReqs 创建包周期存储库请求体
type VaultOrderCreateReqs struct {
	Vault *VaultOrder `json:"vault"`
}

func (o VaultOrderCreateReqs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultOrderCreateReqs struct{}"
	}

	return strings.Join([]string{"VaultOrderCreateReqs", string(data)}, " ")
}
