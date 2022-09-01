package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性ip，同publicips
type EipInfo struct {

	// eip_id
	EipId *string `json:"eip_id,omitempty" xml:"eip_id"`

	// eip_address
	EipAddress *string `json:"eip_address,omitempty" xml:"eip_address"`

	// IP版本号，取值：4表示IPv4,6表示IPv6。  [不支持IPv6，请勿设置为6。](tag:dt,dt_test)
	IpVersion *int32 `json:"ip_version,omitempty" xml:"ip_version"`
}

func (o EipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipInfo struct{}"
	}

	return strings.Join([]string{"EipInfo", string(data)}, " ")
}
