package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgentDaemonsetRequest Request Object
type DeleteAgentDaemonsetRequest struct {

	// Region ID
	Region string `json:"region"`

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群id
	ClusterId string `json:"cluster_id"`

	// 调用服务,默认hss，cce集成防护调用场景使用，包括:    - hss： hss服务    - cce： cce服务，cce集成防护调用需要传参cce
	InvokedService *string `json:"invoked_service,omitempty"`
}

func (o DeleteAgentDaemonsetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgentDaemonsetRequest struct{}"
	}

	return strings.Join([]string{"DeleteAgentDaemonsetRequest", string(data)}, " ")
}
