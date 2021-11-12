package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 域名Id

	InstanceId string `json:"instance_id"`
}

func (o DeleteHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostRequest struct{}"
	}

	return strings.Join([]string{"DeleteHostRequest", string(data)}, " ")
}
