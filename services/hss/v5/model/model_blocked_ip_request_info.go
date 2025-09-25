package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BlockedIpRequestInfo 解除拦截的IP详情
type BlockedIpRequestInfo struct {

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId string `json:"host_id"`

	// 攻击源IP
	SrcIp string `json:"src_ip"`

	// **参数解释**： 登录类型 **约束限制**: 不涉及 **取值范围**: - mysql：mysql服务。 - rdp：rdp服务。 - ssh：ssh服务。 - vsftp：vsftp服务。  **默认取值**: 不涉及
	LoginType string `json:"login_type"`
}

func (o BlockedIpRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlockedIpRequestInfo struct{}"
	}

	return strings.Join([]string{"BlockedIpRequestInfo", string(data)}, " ")
}
