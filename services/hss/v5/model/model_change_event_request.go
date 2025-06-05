package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeEventRequest Request Object
type ChangeEventRequest struct {

	// Region ID
	Region *string `json:"region,omitempty"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 容器实例名称
	ContainerName *string `json:"container_name,omitempty"`

	// 容器Id
	ContainerId *string `json:"container_id,omitempty"`

	Body *ChangeEventRequestInfo `json:"body,omitempty"`
}

func (o ChangeEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEventRequest struct{}"
	}

	return strings.Join([]string{"ChangeEventRequest", string(data)}, " ")
}
