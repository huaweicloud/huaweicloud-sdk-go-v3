package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgentDaemonsetRequest Request Object
type UpdateAgentDaemonsetRequest struct {

	// Region ID
	Region string `json:"region"`

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群id
	ClusterId string `json:"cluster_id"`

	Body *UpdateDaemonsetRequestBody `json:"body,omitempty"`
}

func (o UpdateAgentDaemonsetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgentDaemonsetRequest struct{}"
	}

	return strings.Join([]string{"UpdateAgentDaemonsetRequest", string(data)}, " ")
}
