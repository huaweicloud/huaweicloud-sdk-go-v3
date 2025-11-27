package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// TransitIp 中转子网IP的响应体。
type TransitIp struct {

	// 中转IP的ID。
	Id string `json:"id"`

	// 项目的ID。
	ProjectId string `json:"project_id"`

	// 中转IP的网络接口ID。
	NetworkInterfaceId string `json:"network_interface_id"`

	// 中转IP的地址。
	IpAddress string `json:"ip_address"`

	// 中转IP的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 中转IP的更新时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 当前租户子网的ID。取值约束：与transit_subnet_id参数二选一。默认空字符串。
	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 中转IP绑定的私网NAT网关实例的ID。
	GatewayId string `json:"gateway_id"`

	// 企业项目ID。创建中转IP时，关联的企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 私网NAT中转IP的状态。 取值为： \"ACTIVE\"：正常运行 \"FROZEN\"：冻结 \"INACTIVE\"：不可用
	Status *TransitIpStatus `json:"status,omitempty"`
}

func (o TransitIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransitIp struct{}"
	}

	return strings.Join([]string{"TransitIp", string(data)}, " ")
}

type TransitIpStatus struct {
	value string
}

type TransitIpStatusEnum struct {
	ACTIVE   TransitIpStatus
	FROZEN   TransitIpStatus
	INACTIVE TransitIpStatus
}

func GetTransitIpStatusEnum() TransitIpStatusEnum {
	return TransitIpStatusEnum{
		ACTIVE: TransitIpStatus{
			value: "ACTIVE",
		},
		FROZEN: TransitIpStatus{
			value: "FROZEN",
		},
		INACTIVE: TransitIpStatus{
			value: "INACTIVE",
		},
	}
}

func (c TransitIpStatus) Value() string {
	return c.value
}

func (c TransitIpStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransitIpStatus) UnmarshalJSON(b []byte) error {
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
