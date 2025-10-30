package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentDaemonsetInfoRequest Request Object
type ListAgentDaemonsetInfoRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 集群类型 **约束限制**： 不涉及 **取值范围**： - cce：CCE集群 - ali：阿里云集群 - tencent：腾讯云集群 - azure：微软云集群 - aws：亚马逊集群 - self_built_hw：华为云自建集群 - self_built_idc：IDC自建集群  **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 查询结果是否显示集群日志的接入状态 **约束限制**： 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**： false
	ShowClusterLogStatus *bool `json:"show_cluster_log_status,omitempty"`

	// **参数解释**： 按照集群的接入状态进行查询，true为已接入，false为未接入 **约束限制**： 不涉及 **取值范围**： - true：已接入。 - false：未接入。  **默认取值**： 不涉及
	AccessStatus *bool `json:"access_status,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAgentDaemonsetInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentDaemonsetInfoRequest struct{}"
	}

	return strings.Join([]string{"ListAgentDaemonsetInfoRequest", string(data)}, " ")
}
