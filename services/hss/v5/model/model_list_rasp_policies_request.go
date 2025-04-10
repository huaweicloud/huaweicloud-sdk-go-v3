package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRaspPoliciesRequest Request Object
type ListRaspPoliciesRequest struct {

	// region id
	Region string `json:"region"`

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 操作系统类型，包含如下2种。 - Linux - Windows
	OsType *string `json:"os_type,omitempty"`

	// 每页显示个数，默认10
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`
}

func (o ListRaspPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRaspPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListRaspPoliciesRequest", string(data)}, " ")
}
