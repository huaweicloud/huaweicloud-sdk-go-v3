package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InternalEndpoint struct {

	// 内网访问ID
	Id *string `json:"id,omitempty"`

	// VPC端点的ID
	VpcepEndpointId *string `json:"vpcep_endpoint_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// VPC的ID
	VpcId *string `json:"vpc_id,omitempty"`

	// VPC的名称
	VpcName *string `json:"vpc_name,omitempty"`

	// VPC的CIDR范围
	VpcCidr *string `json:"vpc_cidr,omitempty"`

	// 子网的ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 子网的名称
	SubnetName *string `json:"subnet_name,omitempty"`

	// 子网的CIDR范围
	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	// 端点的IP地址
	EndpointIp *string `json:"endpoint_ip,omitempty"`

	// 访问控制的描述信息
	Description *string `json:"description,omitempty"`

	// 访问控制的状态
	Status *InternalEndpointStatus `json:"status,omitempty"`

	// 访问控制的状态信息
	StatusText *string `json:"status_text,omitempty"`

	// 访问控制的创建时间
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o InternalEndpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InternalEndpoint struct{}"
	}

	return strings.Join([]string{"InternalEndpoint", string(data)}, " ")
}

type InternalEndpointStatus struct {
	value string
}

type InternalEndpointStatusEnum struct {
	CREATING     InternalEndpointStatus
	RUNNING      InternalEndpointStatus
	CREATE_ERROR InternalEndpointStatus
	DELETING     InternalEndpointStatus
	DELETE_ERROR InternalEndpointStatus
}

func GetInternalEndpointStatusEnum() InternalEndpointStatusEnum {
	return InternalEndpointStatusEnum{
		CREATING: InternalEndpointStatus{
			value: "Creating",
		},
		RUNNING: InternalEndpointStatus{
			value: "Running",
		},
		CREATE_ERROR: InternalEndpointStatus{
			value: "CreateError",
		},
		DELETING: InternalEndpointStatus{
			value: "Deleting",
		},
		DELETE_ERROR: InternalEndpointStatus{
			value: "DeleteError",
		},
	}
}

func (c InternalEndpointStatus) Value() string {
	return c.value
}

func (c InternalEndpointStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InternalEndpointStatus) UnmarshalJSON(b []byte) error {
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
