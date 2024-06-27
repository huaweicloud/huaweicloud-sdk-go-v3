package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EipCountRespData EIP 数量查询反馈
type EipCountRespData struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。
	ObjectId *string `json:"object_id,omitempty"`

	// EIP总数
	EipTotal *int32 `json:"eip_total,omitempty"`

	// 该账号下所有墙防护EIP总数量
	EipProtected *int32 `json:"eip_protected,omitempty"`

	// 该当前防火墙防护EIP数量
	EipProtectedSelf *int32 `json:"eip_protected_self,omitempty"`
}

func (o EipCountRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipCountRespData struct{}"
	}

	return strings.Join([]string{"EipCountRespData", string(data)}, " ")
}
