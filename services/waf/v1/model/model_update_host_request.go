package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 域名Id

	InstanceId string `json:"instance_id"`

	Body *UpdateHostRequestBody `json:"body,omitempty"`
}

func (o UpdateHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostRequest", string(data)}, " ")
}
