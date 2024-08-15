package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EipInfo 参数解释：弹性IP，同publicips。
type EipInfo struct {

	// 参数解释：弹性IP的ID。
	EipId *string `json:"eip_id,omitempty"`

	// 参数解释：弹性IP的IP地址。
	EipAddress *string `json:"eip_address,omitempty"`

	// 参数解释：IP版本号。  取值范围： - 4表示IPv4。 - 6表示IPv6。  [不支持IPv6，请勿设置为6。](tag:dt,dt_test)
	IpVersion *int32 `json:"ip_version,omitempty"`
}

func (o EipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipInfo struct{}"
	}

	return strings.Join([]string{"EipInfo", string(data)}, " ")
}
