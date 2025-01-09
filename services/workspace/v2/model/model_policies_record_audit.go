package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesRecordAudit 录屏审计。
type PoliciesRecordAudit struct {

	// 是否开启录屏审计。取值为： false：表示关闭。 true：表示开启。
	Enable *bool `json:"enable,omitempty"`

	Rules *PoliciesRecordAuditRules `json:"rules,omitempty"`
}

func (o PoliciesRecordAudit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesRecordAudit struct{}"
	}

	return strings.Join([]string{"PoliciesRecordAudit", string(data)}, " ")
}
