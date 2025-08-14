package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociatePolicyGroupRequestInfo struct {

	// **参数解释**: 部署的目标策略组ID **约束限制**: 需查询ListPolicyGroup接口，仅支持传其返回参数data_list中support_version等于hss.version.premium或hss.version.container.enterprise的group_id **取值范围**: 只能由英文字母、数字及“-”组成，字符长度36-64位 **默认取值**: 不涉及
	TargetPolicyGroupId string `json:"target_policy_group_id"`

	// **参数解释**: 是否要对全量主机/pod实例/工作负载/集群部署策略，如果为true的话，会自动筛选符合策略组支持版本和操作系统版本的全量主机/pod实例/工作负载/集群部署策略，不需填写host_id_list，如果为false的话，需要填写host_id_list **约束限制**: 不涉及 **取值范围**: - true: 自动筛选符合策略组支持版本和操作系统版本的全量主机/pod实例/工作负载/集群部署策略，不需填写host_id_list - false: 非全量部署，仅对指定的主机/pod实例/工作负载/集群部署策略， 需要填写host_id_list **默认取值**: 不涉及
	OperateAll *bool `json:"operate_all,omitempty"`

	// 策略部署类型: - host: 主机 - pod: pod实例 - workload: 工作负载 - cluster: 集群
	DeployType *string `json:"deploy_type,omitempty"`

	// **参数解释**: 需要部署策略组的主机/pod实例/负载/集群ID列表 **约束限制**: 不涉及 **取值范围**: 最少0条，最多10000条 **默认取值**: 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o AssociatePolicyGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePolicyGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"AssociatePolicyGroupRequestInfo", string(data)}, " ")
}
