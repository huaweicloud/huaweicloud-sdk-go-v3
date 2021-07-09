package model

import (
	"encoding/json"

	"strings"
)

type SetOffSiteBackupPolicyRequestBody struct {
	// 备份策略对象，包括备份类型、备份保留天数、目标区域ID和目标project ID。

	PolicyPara []OffSiteBackupPolicy `json:"policy_para"`
}

func (o SetOffSiteBackupPolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetOffSiteBackupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetOffSiteBackupPolicyRequestBody", string(data)}, " ")
}
