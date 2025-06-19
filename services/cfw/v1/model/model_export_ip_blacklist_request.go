package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportIpBlacklistRequest Request Object
type ExportIpBlacklistRequest struct {

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId string `json:"fw_instance_id"`

	// IP黑名单的名字，如果要导出生效范围为EIP的IP黑名单，就指定名字为ip-blacklist-eip.txt，如果要导出生效范围为NAT的IP黑名单，就指定名字为ip-blacklist-nat.txt。
	Name string `json:"name"`
}

func (o ExportIpBlacklistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportIpBlacklistRequest struct{}"
	}

	return strings.Join([]string{"ExportIpBlacklistRequest", string(data)}, " ")
}
