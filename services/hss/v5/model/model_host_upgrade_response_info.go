package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostUpgradeResponseInfo 主机升级状态信息
type HostUpgradeResponseInfo struct {

	// **参数解释**: 服务器ID **取值范围**: 字符长度0-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 升级状态 **取值范围**: - not_upgrade：未升级，默认状态，客户还没有给这台机器下发过升级 - upgrading：正在升级到2.0中 - upgrade_failed：升级到2.0失败 - upgrade_succeed：已经成功升级到2.0
	UpgradeStatus *string `json:"upgrade_status,omitempty"`
}

func (o HostUpgradeResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostUpgradeResponseInfo struct{}"
	}

	return strings.Join([]string{"HostUpgradeResponseInfo", string(data)}, " ")
}
