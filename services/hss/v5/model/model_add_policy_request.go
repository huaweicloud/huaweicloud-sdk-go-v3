package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPolicyRequest Request Object
type AddPolicyRequest struct {

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 操作系统类型，包含如下2种。 - Linux - Windows
	OsType *string `json:"os_type,omitempty"`

	// 策略名称
	PolicyName string `json:"policy_name"`

	Body *AddPolicyRequestInfo `json:"body,omitempty"`
}

func (o AddPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPolicyRequest struct{}"
	}

	return strings.Join([]string{"AddPolicyRequest", string(data)}, " ")
}
