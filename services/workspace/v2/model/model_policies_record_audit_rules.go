package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesRecordAuditRules 录屏规则列表。
type PoliciesRecordAuditRules struct {
	Rule *RecordAuditRule `json:"rule,omitempty"`
}

func (o PoliciesRecordAuditRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesRecordAuditRules struct{}"
	}

	return strings.Join([]string{"PoliciesRecordAuditRules", string(data)}, " ")
}
