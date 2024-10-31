package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeEastWestFirewallStatusResponseData struct {

	// 东西向防护对象ID，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。此处仅取type为1的防护对象id，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。
	Id *string `json:"id,omitempty"`
}

func (o ChangeEastWestFirewallStatusResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEastWestFirewallStatusResponseData struct{}"
	}

	return strings.Join([]string{"ChangeEastWestFirewallStatusResponseData", string(data)}, " ")
}
