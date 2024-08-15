package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GlobalEipInfo struct {

	// 参数解释：global eip的id
	GlobalEipId *string `json:"global_eip_id,omitempty"`

	// 参数解释：global eip的ip地址
	GlobalEipAddress *string `json:"global_eip_address,omitempty"`

	// 参数解释：IP版本信息。  取值范围： - 4，表示IPv4 - 6，表示IPv6  [不支持IPv6，请勿设置为6。](tag:dt,dt_test)
	IpVersion *int32 `json:"ip_version,omitempty"`
}

func (o GlobalEipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalEipInfo struct{}"
	}

	return strings.Join([]string{"GlobalEipInfo", string(data)}, " ")
}
