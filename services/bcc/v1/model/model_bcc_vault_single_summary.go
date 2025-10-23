package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BccVaultSingleSummary 单个存储库summary返回body
type BccVaultSingleSummary struct {

	// vault的类型
	VaultObjectType *string `json:"vault_object_type,omitempty"`

	// vault的数量
	VaultCount *int64 `json:"vault_count,omitempty"`

	// vault的大小
	VaultSumSize *int64 `json:"vault_sum_size,omitempty"`

	// vault的使用大小
	VaultUsedSize *int64 `json:"vault_used_size,omitempty"`

	// vault的使用率
	VaultUtilizationRate *float64 `json:"vault_utilization_rate,omitempty"`
}

func (o BccVaultSingleSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BccVaultSingleSummary struct{}"
	}

	return strings.Join([]string{"BccVaultSingleSummary", string(data)}, " ")
}
