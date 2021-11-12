package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 终端节点服务列表
type EndpointServiceS struct {
	// 公共终端节点服务的ID，唯一标 识。

	Id *string `json:"id,omitempty"`
	// 终端节点服务的所有者。

	Owner *string `json:"owner,omitempty"`
	// 公共终端节点服务的名称。

	ServiceName *string `json:"service_name,omitempty"`
	// 终端节点服务类型。 ● gataway：由运维人员配置。用 户无需创建，可直接使用。 ● interface：包括运维人员配置的 云服务和用户自己创建的私有服 务。其中，运维人员配置的云服 务无需创建，用户可直接使用。 您可以通过创建终端节点创建访问 Gateway和Interface类型终端节点 服务的终端节点。

	ServiceType *EndpointServiceSServiceType `json:"service_type,omitempty"`
	// 终端节点服务的创建时间。 采用UTC时间格式，格式为：YYYYMM- DDTHH:MM:SSZ

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
	// 连接该终端节点服务的终端节点是 否计费。 ● true：计费 ● false：不计费

	IsCharge *bool `json:"is_charge,omitempty"`
}

func (o EndpointServiceS) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointServiceS struct{}"
	}

	return strings.Join([]string{"EndpointServiceS", string(data)}, " ")
}

type EndpointServiceSServiceType struct {
	value string
}

type EndpointServiceSServiceTypeEnum struct {
	INTERFACE EndpointServiceSServiceType
	GATEWAY   EndpointServiceSServiceType
}

func GetEndpointServiceSServiceTypeEnum() EndpointServiceSServiceTypeEnum {
	return EndpointServiceSServiceTypeEnum{
		INTERFACE: EndpointServiceSServiceType{
			value: "interface",
		},
		GATEWAY: EndpointServiceSServiceType{
			value: "gateway",
		},
	}
}

func (c EndpointServiceSServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointServiceSServiceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
