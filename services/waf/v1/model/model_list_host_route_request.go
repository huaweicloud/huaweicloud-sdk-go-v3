package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHostRouteRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 域名Id

	InstanceId string `json:"instance_id"`
}

func (o ListHostRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostRouteRequest struct{}"
	}

	return strings.Join([]string{"ListHostRouteRequest", string(data)}, " ")
}
