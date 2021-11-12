package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePremiumHostProtectStatusRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 独享模式域名ID

	HostId string `json:"host_id"`

	Body *UpdatePremiumHostProtectStatusRequestBody `json:"body,omitempty"`
}

func (o UpdatePremiumHostProtectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostProtectStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostProtectStatusRequest", string(data)}, " ")
}
