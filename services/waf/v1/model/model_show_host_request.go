package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 域名Id

	InstanceId string `json:"instance_id"`
}

func (o ShowHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostRequest struct{}"
	}

	return strings.Join([]string{"ShowHostRequest", string(data)}, " ")
}
