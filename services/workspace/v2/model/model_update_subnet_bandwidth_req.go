package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSubnetBandwidthReq struct {

	// 云办公带宽名称。
	BandwidthName *string `json:"bandwidth_name,omitempty"`

	// 云办公带宽大小。
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 企业项目ID，默认\"0\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 包周期订购ID，CBC订购回调时使用。
	OrderId *string `json:"order_id,omitempty"`

	// 计费模式 - wks_bandwidth：按带宽计费。 - free: 不计费，不支持包周期订购。 - wks_traffic：按流量计费
	ChargeMode *string `json:"charge_mode,omitempty"`
}

func (o UpdateSubnetBandwidthReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetBandwidthReq struct{}"
	}

	return strings.Join([]string{"UpdateSubnetBandwidthReq", string(data)}, " ")
}
