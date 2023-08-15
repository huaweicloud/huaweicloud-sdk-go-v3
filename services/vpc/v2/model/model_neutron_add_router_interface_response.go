package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NeutronAddRouterInterfaceResponse Response Object
type NeutronAddRouterInterfaceResponse struct {

	// 路由器添加的端口对应的ID
	PortId *string `json:"port_id,omitempty"`

	// 路由器添加的子网对应的ID
	SubnetId *NeutronAddRouterInterfaceResponseSubnetId `json:"subnet_id,omitempty"`

	// 项目ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 路由器的ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o NeutronAddRouterInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronAddRouterInterfaceResponse struct{}"
	}

	return strings.Join([]string{"NeutronAddRouterInterfaceResponse", string(data)}, " ")
}

type NeutronAddRouterInterfaceResponseSubnetId struct {
	value string
}

type NeutronAddRouterInterfaceResponseSubnetIdEnum struct {
	E_0_9A_F8__0_9A_F4__0_9A_F4__0_9A_F4__0_9A_F12 NeutronAddRouterInterfaceResponseSubnetId
}

func GetNeutronAddRouterInterfaceResponseSubnetIdEnum() NeutronAddRouterInterfaceResponseSubnetIdEnum {
	return NeutronAddRouterInterfaceResponseSubnetIdEnum{
		E_0_9A_F8__0_9A_F4__0_9A_F4__0_9A_F4__0_9A_F12: NeutronAddRouterInterfaceResponseSubnetId{
			value: "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		},
	}
}

func (c NeutronAddRouterInterfaceResponseSubnetId) Value() string {
	return c.value
}

func (c NeutronAddRouterInterfaceResponseSubnetId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronAddRouterInterfaceResponseSubnetId) UnmarshalJSON(b []byte) error {
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
