package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErInstance struct {

	// Er实例id
	Id *string `json:"id,omitempty"`

	// Er名称
	Name *string `json:"name,omitempty"`

	// Er状态
	State *string `json:"state,omitempty"`

	// 企业租户id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 租户id
	ProjectId *string `json:"project_id,omitempty"`

	// 是否开启ipv6
	EnableIpv6 *string `json:"enable_ipv6,omitempty"`
}

func (o ErInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErInstance struct{}"
	}

	return strings.Join([]string{"ErInstance", string(data)}, " ")
}
