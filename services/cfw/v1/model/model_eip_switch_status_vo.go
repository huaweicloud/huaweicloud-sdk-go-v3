package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EipSwitchStatusVo struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID。此处仅取type为0的防护对象id，可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得
	ObjectId *string `json:"object_id,omitempty"`

	// 修改eip防护状态失败状态列表，状态包括成功“successful”，失败“fail”
	FailEipIdList *[]string `json:"fail_eip_id_list,omitempty"`

	// 修改eip防护状态失败信息列表
	FailEipList *[]FailedEipInfo `json:"fail_eip_list,omitempty"`

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	Id *string `json:"id,omitempty"`
}

func (o EipSwitchStatusVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipSwitchStatusVo struct{}"
	}

	return strings.Join([]string{"EipSwitchStatusVo", string(data)}, " ")
}
