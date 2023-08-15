package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EndpointService 终端节点服务列表
type EndpointService struct {

	// 公共终端节点服务的ID，唯一标识。
	Id *string `json:"id,omitempty"`

	// 终端节点服务的所有者。
	Owner *string `json:"owner,omitempty"`

	// 公共终端节点服务的名称。
	ServiceName *string `json:"service_name,omitempty"`

	// 终端节点服务类型。 ● gataway：由运维人员配置。用户无需创建，可直接使用。 ● interface：包括运维人员配置的云服务和用户自己创建的私有服务。 其中，运维人员配置的云服务无需创建， 用户可直接使用。 您可以通过创建终端节点创建访问Gateway和Interface类型终端节点服务的终端节点。
	ServiceType *EndpointServiceServiceType `json:"service_type,omitempty"`

	// 终端节点服务的创建时间。 采用UTC时间格式，格式为：YYYYMM-DDTHH:MM:SSZ
	CreatedAt *string `json:"created_at,omitempty"`

	// 连接该终端节点服务的终端节点是否计费。 ● true：计费 ● false：不计费
	IsCharge *bool `json:"is_charge,omitempty"`

	// 是否开启终端节点策略。 ● false：不支持设置终端节点策略 ● true：支持设置终端节点策略 默认为false 是否开启终端节点策略。 ● false：不支持设置终端节点策略 ● true：支持设置终端节点策略 默认为false
	EnablePolicy *bool `json:"enable_policy,omitempty"`
}

func (o EndpointService) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointService struct{}"
	}

	return strings.Join([]string{"EndpointService", string(data)}, " ")
}

type EndpointServiceServiceType struct {
	value string
}

type EndpointServiceServiceTypeEnum struct {
	INTERFACE EndpointServiceServiceType
	GATEWAY   EndpointServiceServiceType
}

func GetEndpointServiceServiceTypeEnum() EndpointServiceServiceTypeEnum {
	return EndpointServiceServiceTypeEnum{
		INTERFACE: EndpointServiceServiceType{
			value: "interface",
		},
		GATEWAY: EndpointServiceServiceType{
			value: "gateway",
		},
	}
}

func (c EndpointServiceServiceType) Value() string {
	return c.value
}

func (c EndpointServiceServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointServiceServiceType) UnmarshalJSON(b []byte) error {
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
