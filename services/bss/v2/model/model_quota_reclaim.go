package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type QuotaReclaim struct {

	// 被回收的云经销商的代金券额度ID。
	QuotaId *string `json:"quota_id,omitempty"`

	// 被回收额度后的代金券额度余额。单位：元。
	QuotaBalance *decimal.Decimal `json:"quota_balance,omitempty"`
}

func (o QuotaReclaim) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaReclaim struct{}"
	}

	return strings.Join([]string{"QuotaReclaim", string(data)}, " ")
}
