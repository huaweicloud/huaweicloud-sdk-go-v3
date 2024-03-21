package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigurationDataItems 访问方式配置项。
type AccessConfigurationDataItems struct {

	// 访问方式的uid。
	Uid *string `json:"uid,omitempty"`

	Metadata *AccessConfigurationMetadata `json:"metadata,omitempty"`

	// 访问方式类型。
	Type *AccessConfigurationDataItemsType `json:"type,omitempty"`

	// 内网访问方式域名。
	DomainNames *[]string `json:"domain_names,omitempty"`

	AccessControl *AccessControl `json:"access_control,omitempty"`

	// 访问方式配置端口、协议、证书、URL路径等信息列表。
	Ports *[]AccessConfigurationPort `json:"ports,omitempty"`

	// 用户选择的elb的ID。
	ElbId *string `json:"elb_id,omitempty"`

	// 响应体参数，用户选择的elb的公网ip。
	PublicIp *string `json:"public_ip,omitempty"`

	// 响应体参数，用户选择的elb的私网ip。
	PrivateIp *string `json:"private_ip,omitempty"`
}

func (o AccessConfigurationDataItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigurationDataItems struct{}"
	}

	return strings.Join([]string{"AccessConfigurationDataItems", string(data)}, " ")
}

type AccessConfigurationDataItemsType struct {
	value string
}

type AccessConfigurationDataItemsTypeEnum struct {
	CLUSTER_IP    AccessConfigurationDataItemsType
	LOAD_BALANCER AccessConfigurationDataItemsType
	INGRESS       AccessConfigurationDataItemsType
}

func GetAccessConfigurationDataItemsTypeEnum() AccessConfigurationDataItemsTypeEnum {
	return AccessConfigurationDataItemsTypeEnum{
		CLUSTER_IP: AccessConfigurationDataItemsType{
			value: "ClusterIP",
		},
		LOAD_BALANCER: AccessConfigurationDataItemsType{
			value: "LoadBalancer",
		},
		INGRESS: AccessConfigurationDataItemsType{
			value: "Ingress",
		},
	}
}

func (c AccessConfigurationDataItemsType) Value() string {
	return c.value
}

func (c AccessConfigurationDataItemsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigurationDataItemsType) UnmarshalJSON(b []byte) error {
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
