package model

import (
	"encoding/json"

	"strings"
)

type TenantConfigV2 struct {
	// 抄送邮箱最大个数

	CreateCaseCcemailMaxCount *int32 `json:"create_case_ccemail_max_count,omitempty"`
}

func (o TenantConfigV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TenantConfigV2 struct{}"
	}

	return strings.Join([]string{"TenantConfigV2", string(data)}, " ")
}
