package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VaultCreateParameters 存储库创建参数
type VaultCreateParameters struct {
	CombinedOrder *CombinedOrder `json:"combined_order,omitempty"`
}

func (o VaultCreateParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultCreateParameters struct{}"
	}

	return strings.Join([]string{"VaultCreateParameters", string(data)}, " ")
}
