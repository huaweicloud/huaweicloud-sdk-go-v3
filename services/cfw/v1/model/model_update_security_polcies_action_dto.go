package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecurityPolciesActionDto struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。
	ObjectId string `json:"object_id"`

	// 动作
	Action string `json:"action"`

	// 规则ID列表
	RuleIds []string `json:"rule_ids"`
}

func (o UpdateSecurityPolciesActionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolciesActionDto struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolciesActionDto", string(data)}, " ")
}
