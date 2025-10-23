package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BccVaultDaySummary 存储库summary返回body
type BccVaultDaySummary struct {

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 时间戳
	TimeStamp *int64 `json:"time_stamp,omitempty"`

	// vault的summary列表
	VaultSummaryList *[]BccVaultSingleSummary `json:"vault_summary_list,omitempty"`

	// 存储库summary数量
	VaultSummaryCount *int64 `json:"vault_summary_count,omitempty"`

	// 存储库summary全部容量大小，单位GB。
	VaultSummarySize *int64 `json:"vault_summary_size,omitempty"`

	// 存储库summary已用容量大小，单位MB。
	VaultSummaryUsed *int64 `json:"vault_summary_used,omitempty"`

	// 存储库summary容量使用率
	VaultAverageUtilizationRate *float64 `json:"vault_average_utilization_rate,omitempty"`
}

func (o BccVaultDaySummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BccVaultDaySummary struct{}"
	}

	return strings.Join([]string{"BccVaultDaySummary", string(data)}, " ")
}
