package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetwork 中心网络。
type CreateCentralNetwork struct {

	// 实例名称。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例标签。
	Tags *[]Tag `json:"tags,omitempty"`

	// 实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	PolicyDocument *CentralNetworkPolicyDocument `json:"policy_document,omitempty"`
}

func (o CreateCentralNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetwork struct{}"
	}

	return strings.Join([]string{"CreateCentralNetwork", string(data)}, " ")
}
