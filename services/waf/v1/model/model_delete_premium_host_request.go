package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePremiumHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 独享模式域名ID

	HostId string `json:"host_id"`
	// 是否保留规则

	KeepPolicy *bool `json:"keepPolicy,omitempty"`
}

func (o DeletePremiumHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePremiumHostRequest struct{}"
	}

	return strings.Join([]string{"DeletePremiumHostRequest", string(data)}, " ")
}
