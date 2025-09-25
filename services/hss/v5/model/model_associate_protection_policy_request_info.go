package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateProtectionPolicyRequestInfo struct {

	// 部署的目标策略组ID
	TargetProtectionPolicyId string `json:"target_protection_policy_id"`

	// **参数解释**: 服务器ID列表 **取值范围**: 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o AssociateProtectionPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateProtectionPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"AssociateProtectionPolicyRequestInfo", string(data)}, " ")
}
