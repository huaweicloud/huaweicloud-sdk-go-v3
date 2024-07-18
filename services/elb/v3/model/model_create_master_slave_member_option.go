package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateMasterSlaveMemberOption 创建主备后端服务器请求参数
type CreateMasterSlaveMemberOption struct {

	// 后端服务器对应的IP地址。  使用说明： - 若subnet_cidr_id为空，表示添加跨VPC后端，此时address必须为IPv4地址。 - 若subnet_cidr_id不为空，表示是一个关联到ECS的后端服务器。 该IP地址可以是私网IPv4或IPv6。但必须在subnet_cidr_id对应的子网网段中。且只能指定为关联ECS的主网卡内网IP。  [不支持IPv6，请勿设置为IPv6地址。](tag:dt,dt_test)
	Address string `json:"address"`

	// 后端云服务器的管理状态。  取值：true。  虽然创建、更新请求支持该字段，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 后端云服务器名称。
	Name *string `json:"name,omitempty"`

	// 后端服务器业务端口。 >在开启端口透传的pool下创建member传该字段不生效，可不传该字段。
	ProtocolPort int32 `json:"protocol_port"`

	// 后端云服务器所在的子网ID，可以是子网的IPv4子网ID或IPv6子网ID。  使用说明： - 该子网和关联的负载均衡器的子网必须在同一VPC下。 - 若所属LB的跨VPC后端转发特性已开启，则该字段可以不传，表示添加跨VPC的后端服务器。 此时address必须为IPv4地址，所在的pool的协议必须为TCP/HTTP/HTTPS。  [不支持IPv6，请勿设置为IPv6子网ID。](tag:dt,dt_test)
	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`

	// 后端服务器的主备状态。  取值范围： - master：主后端服务器。 - slave：备后端服务器。
	Role CreateMasterSlaveMemberOptionRole `json:"role"`
}

func (o CreateMasterSlaveMemberOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMasterSlaveMemberOption struct{}"
	}

	return strings.Join([]string{"CreateMasterSlaveMemberOption", string(data)}, " ")
}

type CreateMasterSlaveMemberOptionRole struct {
	value string
}

type CreateMasterSlaveMemberOptionRoleEnum struct {
	MASTER CreateMasterSlaveMemberOptionRole
	SLAVE  CreateMasterSlaveMemberOptionRole
}

func GetCreateMasterSlaveMemberOptionRoleEnum() CreateMasterSlaveMemberOptionRoleEnum {
	return CreateMasterSlaveMemberOptionRoleEnum{
		MASTER: CreateMasterSlaveMemberOptionRole{
			value: "master",
		},
		SLAVE: CreateMasterSlaveMemberOptionRole{
			value: "slave",
		},
	}
}

func (c CreateMasterSlaveMemberOptionRole) Value() string {
	return c.value
}

func (c CreateMasterSlaveMemberOptionRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMasterSlaveMemberOptionRole) UnmarshalJSON(b []byte) error {
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
