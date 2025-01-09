package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplySubnetBandwidthReq struct {

	// 云办公带宽名称。
	BandwidthName *string `json:"bandwidth_name,omitempty"`

	// 子网id。
	SubnetId string `json:"subnet_id"`

	// 计费模式 - wks_bandwidth：按带宽计费，仅包周期支持。 - free: 不计费，仅按需支持。 - wks_traffic：按流量计费，仅按需支持。
	ChargeMode string `json:"charge_mode"`

	// 云办公带宽大小。
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 企业项目ID，默认\"0\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 包周期订购ID，CBC订购回调时使用。
	OrderId *string `json:"order_id,omitempty"`
}

func (o ApplySubnetBandwidthReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplySubnetBandwidthReq struct{}"
	}

	return strings.Join([]string{"ApplySubnetBandwidthReq", string(data)}, " ")
}
