package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityGroupPolicyRequest Request Object
type UpdateSecurityGroupPolicyRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** Kubernetes集群的唯一标识，用于指定更新安全组策略所属的集群 **约束限制** 需确保集群已接入HSS服务，且账号拥有该集群的操作权限 **取值范围** 字符长度1-64位，支持字母、数字、短横线（-）、下划线（_） **默认取值** 无
	ClusterId string `json:"cluster_id"`

	// **参数解释** Kubernetes集群内的命名空间标识，用于隔离不同命名空间下的安全组策略 **约束限制** 命名空间需已存在于指定集群中，否则返回资源不存在错误 **取值范围** 字符长度1-63位，支持字母、数字、短横线（-），且不能以短横线开头或结尾 **默认取值** 无
	Namespace string `json:"namespace"`

	Body *UpdateSecurityGroupPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateSecurityGroupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupPolicyRequest", string(data)}, " ")
}
