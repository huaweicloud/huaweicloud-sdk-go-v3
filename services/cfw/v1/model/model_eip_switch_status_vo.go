package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EipSwitchStatusVo struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。
	ObjectId *string `json:"object_id,omitempty"`

	// 修改eip防护状态失败列表。
	FailEipIdList *[]string `json:"fail_eip_id_list,omitempty"`

	// ID
	Id *string `json:"id,omitempty"`
}

func (o EipSwitchStatusVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipSwitchStatusVo struct{}"
	}

	return strings.Join([]string{"EipSwitchStatusVo", string(data)}, " ")
}
