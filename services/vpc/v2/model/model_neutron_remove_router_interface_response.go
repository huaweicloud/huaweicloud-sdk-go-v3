package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NeutronRemoveRouterInterfaceResponse Response Object
type NeutronRemoveRouterInterfaceResponse struct {

	// 路由器添加的端口对应的ID
	PortId *string `json:"port_id,omitempty"`

	// 路由器添加的子网对应的ID
	SubnetId *NeutronRemoveRouterInterfaceResponseSubnetId `json:"subnet_id,omitempty"`

	// 项目ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 路由器的ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o NeutronRemoveRouterInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronRemoveRouterInterfaceResponse struct{}"
	}

	return strings.Join([]string{"NeutronRemoveRouterInterfaceResponse", string(data)}, " ")
}

type NeutronRemoveRouterInterfaceResponseSubnetId struct {
	value string
}

type NeutronRemoveRouterInterfaceResponseSubnetIdEnum struct {
	E_0_9A_F8__0_9A_F4__0_9A_F4__0_9A_F4__0_9A_F12 NeutronRemoveRouterInterfaceResponseSubnetId
}

func GetNeutronRemoveRouterInterfaceResponseSubnetIdEnum() NeutronRemoveRouterInterfaceResponseSubnetIdEnum {
	return NeutronRemoveRouterInterfaceResponseSubnetIdEnum{
		E_0_9A_F8__0_9A_F4__0_9A_F4__0_9A_F4__0_9A_F12: NeutronRemoveRouterInterfaceResponseSubnetId{
			value: "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		},
	}
}

func (c NeutronRemoveRouterInterfaceResponseSubnetId) Value() string {
	return c.value
}

func (c NeutronRemoveRouterInterfaceResponseSubnetId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronRemoveRouterInterfaceResponseSubnetId) UnmarshalJSON(b []byte) error {
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
