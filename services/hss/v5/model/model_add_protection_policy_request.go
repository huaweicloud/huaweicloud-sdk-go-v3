package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddProtectionPolicyRequest Request Object
type AddProtectionPolicyRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// region id
	Region *string `json:"region,omitempty"`

	Body *AddProtectionPolicyInfoRequestInfo `json:"body,omitempty"`
}

func (o AddProtectionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectionPolicyRequest struct{}"
	}

	return strings.Join([]string{"AddProtectionPolicyRequest", string(data)}, " ")
}
