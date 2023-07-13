package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindedEipResult 查询已绑定EIP记录明细信息。
type BindedEipResult struct {

	// 弹性公网ID。
	PublicIpId string `json:"public_ip_id"`

	// 弹性公网类型。
	PublicIpType string `json:"public_ip_type"`

	// 端口ID。
	PortId string `json:"port_id"`

	// 弹性公网IP。
	PublicIpAddress string `json:"public_ip_address"`

	// 内网地址。
	PrivateIpAddress string `json:"private_ip_address"`

	// 带宽ID。
	BandwidthId string `json:"bandwidth_id"`

	// 带宽名称。
	BandwidthName string `json:"bandwidth_name"`

	// 带宽共享类型。
	BandwidthShareType string `json:"bandwidth_share_type"`

	// 带宽大小。
	BandwidthSize int32 `json:"bandwidth_size"`

	// 修改时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+08:00。
	AppliedAt string `json:"applied_at"`
}

func (o BindedEipResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindedEipResult struct{}"
	}

	return strings.Join([]string{"BindedEipResult", string(data)}, " ")
}
