package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ContainerNetwork Container network parameters.
type ContainerNetwork struct {

	// 容器网络类型（只可选择其一） - overlay_l2：容器隧道网络，通过OVS（OpenVSwitch）为容器构建的overlay_l2网络。 - vpc-router：VPC网络，使用ipvlan和自定义VPC路由为容器构建的Underlay的l2网络。 [- eni：云原生网络2.0，深度整合VPC原生ENI弹性网卡能力，采用VPC网段分配容器地址，支持ELB直通容器，享有高性能，创建CCE Turbo集群时指定。](tag:hws,hws_hk,dt)
	Mode ContainerNetworkMode `json:"mode"`

	// 容器网络网段，建议使用网段10.0.0.0/12~19，172.16.0.0/16~19，192.168.0.0/16~19，如存在网段冲突，将会报错。  此参数在集群创建后不可更改，请谨慎选择。（已废弃，如填写cidrs将忽略该cidr）
	Cidr *string `json:"cidr,omitempty"`

	// 容器网络网段列表。1.21及新版本集群使用cidrs字段，当集群网络类型为vpc-router类型时，支持多个容器网段；1.21之前版本若使用cidrs字段，则取值cidrs数组中的第一个cidr元素作为容器网络网段地址。  此参数在集群创建后不可更改，请谨慎选择。
	Cidrs *[]ContainerCidr `json:"cidrs,omitempty"`
}

func (o ContainerNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerNetwork struct{}"
	}

	return strings.Join([]string{"ContainerNetwork", string(data)}, " ")
}

type ContainerNetworkMode struct {
	value string
}

type ContainerNetworkModeEnum struct {
	OVERLAY_L2 ContainerNetworkMode
	VPC_ROUTER ContainerNetworkMode
	ENI        ContainerNetworkMode
}

func GetContainerNetworkModeEnum() ContainerNetworkModeEnum {
	return ContainerNetworkModeEnum{
		OVERLAY_L2: ContainerNetworkMode{
			value: "overlay_l2",
		},
		VPC_ROUTER: ContainerNetworkMode{
			value: "vpc-router",
		},
		ENI: ContainerNetworkMode{
			value: "eni",
		},
	}
}

func (c ContainerNetworkMode) Value() string {
	return c.value
}

func (c ContainerNetworkMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContainerNetworkMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
