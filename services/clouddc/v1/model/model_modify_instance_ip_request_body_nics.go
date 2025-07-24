package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyInstanceIpRequestBodyNics struct {

	// 创建网卡所属 CloudDCN 子网 ID，可通过 CloudDCN API 查询：https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html。
	SubnetId string `json:"subnet_id"`

	// 创建网卡指定静态IP，批量安装时，该参数不生效。不填或空字符串，默认在子网（subnet）中自动分配一个未使用的IP作网卡的IP地址。若指定IP地址，该IP地址必须在子网（subnet）对应的网段内，且未被使用。
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o ModifyInstanceIpRequestBodyNics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceIpRequestBodyNics struct{}"
	}

	return strings.Join([]string{"ModifyInstanceIpRequestBodyNics", string(data)}, " ")
}
