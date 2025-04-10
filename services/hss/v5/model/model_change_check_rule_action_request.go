package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCheckRuleActionRequest Request Object
type ChangeCheckRuleActionRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 主机ID，不赋值时，查租户所有主机
	HostId *string `json:"host_id,omitempty"`

	// 是否校验cce
	CheckCce *bool `json:"check_cce,omitempty"`

	// 动作 - \"ignore\" - \"unignore\" - \"fix\" - \"verify\"
	Action string `json:"action"`

	Body *CheckRuleIdListRequestInfo `json:"body,omitempty"`
}

func (o ChangeCheckRuleActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCheckRuleActionRequest struct{}"
	}

	return strings.Join([]string{"ChangeCheckRuleActionRequest", string(data)}, " ")
}
