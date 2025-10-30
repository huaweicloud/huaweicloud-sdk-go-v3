package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulRepairCmdsRequestBody **参数解释**: 漏洞操作信息 **取值范围**: 不涉及
type ListVulRepairCmdsRequestBody struct {

	// **参数解释**: 选择全部漏洞类型,该参数优先级高于data_list和host_data_list **约束限制**: 不涉及 **取值范围**: - all_vul：选择全部漏洞 - all_host：选择全部主机漏洞 - all_vul_container_host：选择所有容器节点 - all_vul_container：选择所有容器  **默认取值**: 不涉及
	SelectType *string `json:"select_type,omitempty"`

	// **参数解释**: 漏洞列表 **约束限制**: 不涉及 **取值范围**: 最少1条，最多500条 **默认取值**: 不涉及
	DataList *[]VulOperateInfo `json:"data_list,omitempty"`

	// **参数解释**: 主机维度漏洞列表 **约束限制**: 不涉及 **取值范围**: 最少1条，最多500条 **默认取值**: 不涉及
	HostDataList *[]HostVulOperateInfo `json:"host_data_list,omitempty"`

	// **参数解释**: 是否是容器场景 **约束限制**: 不涉及 **取值范围**: - true：是容器场景 - false：不是容器场景 **默认取值**: false
	IsContainer *bool `json:"is_container,omitempty"`

	// **参数解释**: 集群id **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ListVulRepairCmdsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulRepairCmdsRequestBody struct{}"
	}

	return strings.Join([]string{"ListVulRepairCmdsRequestBody", string(data)}, " ")
}
