package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVaultSummaryResponse Response Object
type ShowVaultSummaryResponse struct {

	// 列举vault的summary
	ListVaultSummary *[]BccVaultDaySummary `json:"list_vault_summary,omitempty"`

	// 存储库数量
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVaultSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowVaultSummaryResponse", string(data)}, " ")
}
