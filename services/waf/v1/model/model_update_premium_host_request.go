package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePremiumHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 独享模式域名ID

	HostId string `json:"host_id"`

	Body *UpdatePremiumHostRequestBody `json:"body,omitempty"`
}

func (o UpdatePremiumHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostRequest struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostRequest", string(data)}, " ")
}
