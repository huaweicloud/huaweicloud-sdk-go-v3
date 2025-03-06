package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NicForSingleServer struct {

	// 租户自定义的子网 ID，为待创建的云手机裸服务器所属的子网。 需要指定vpc_id对应VPC下已创建的子网（subnet）的网络ID，UUID格式
	SubnetId string `json:"subnet_id"`

	// 是否支持ipv6，默认不支持ipv6。 false：不支持ipv6 true：支持ipv6
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// ipv6绑定的共享带宽ID。
	Ipv6BandWidthId *string `json:"ipv6_band_width_id,omitempty"`
}

func (o NicForSingleServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NicForSingleServer struct{}"
	}

	return strings.Join([]string{"NicForSingleServer", string(data)}, " ")
}
