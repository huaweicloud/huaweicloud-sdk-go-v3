package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyInternetReq struct {

	// 公网NAT网关的规格，1：小型，2：中型，3：大型，4：超大型。
	NatSpec string `json:"nat_spec"`

	// traffic（按流量计费），bandwidth（按带宽计费）。
	EipChargeMode string `json:"eip_charge_mode"`

	// 带宽大小。
	BandwidthSize int32 `json:"bandwidth_size"`

	// EIP的类型，5_bgp（全动态BGP），5_sbgp（静态BGP），默认值：5_bgp。
	EipType *string `json:"eip_type,omitempty"`

	// vpc的id。
	VpcId string `json:"vpc_id"`

	// 子网的id。
	SubnetId string `json:"subnet_id"`

	// 企业项目ID，默认\"0。\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// nat的id，有传则使用该NAT，否则新建。
	NatId *string `json:"nat_id,omitempty"`

	// nat名称，默认值：nat-workspace。
	NatName *string `json:"nat_name,omitempty"`

	// eip名称，默认值：eip-workspace。
	EipName *string `json:"eip_name,omitempty"`
}

func (o ApplyInternetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyInternetReq struct{}"
	}

	return strings.Join([]string{"ApplyInternetReq", string(data)}, " ")
}
