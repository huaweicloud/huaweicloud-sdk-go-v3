package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpsProtectDto description
type IpsProtectDto struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。
	ObjectId *string `json:"object_id,omitempty"`

	// ips防护模式，0：观察模式，1：严格模式，2：中等模式，3：宽松模式
	Mode *int32 `json:"mode,omitempty"`
}

func (o IpsProtectDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsProtectDto struct{}"
	}

	return strings.Join([]string{"IpsProtectDto", string(data)}, " ")
}
