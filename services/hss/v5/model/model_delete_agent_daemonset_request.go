package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgentDaemonsetRequest Request Object
type DeleteAgentDaemonsetRequest struct {

	// Region ID
	Region string `json:"region"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
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
