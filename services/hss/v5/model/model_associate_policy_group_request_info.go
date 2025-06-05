package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociatePolicyGroupRequestInfo struct {

	// 部署的目标策略组ID
	TargetPolicyGroupId string `json:"target_policy_group_id"`

	// 是否要对全量主机/pod实例/工作负载/集群部署策略，如果为true的话，不需填写host_id_list，如果为false的话，需要填写host_id_list
	OperateAll *bool `json:"operate_all,omitempty"`

	// 策略部署类型: - host: 主机 - pod: pod实例 - workload: 工作负载 - cluster: 集群
	DeployType *string `json:"deploy_type,omitempty"`

	// 需要部署策略组的主机/pod实例/负载/集群ID列表
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o AssociatePolicyGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePolicyGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"AssociatePolicyGroupRequestInfo", string(data)}, " ")
}
