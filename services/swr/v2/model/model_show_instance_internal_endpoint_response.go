package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowInstanceInternalEndpointResponse Response Object
type ShowInstanceInternalEndpointResponse struct {

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
	Status *ShowInstanceInternalEndpointResponseStatus `json:"status,omitempty"`

	// 访问控制的状态信息
	StatusText *string `json:"status_text,omitempty"`

	// 访问控制的创建时间
	CreatedAt      *sdktime.SdkTime `json:"created_at,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowInstanceInternalEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceInternalEndpointResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceInternalEndpointResponse", string(data)}, " ")
}

type ShowInstanceInternalEndpointResponseStatus struct {
	value string
}

type ShowInstanceInternalEndpointResponseStatusEnum struct {
	CREATING     ShowInstanceInternalEndpointResponseStatus
	RUNNING      ShowInstanceInternalEndpointResponseStatus
	CREATE_ERROR ShowInstanceInternalEndpointResponseStatus
	DELETING     ShowInstanceInternalEndpointResponseStatus
	DELETE_ERROR ShowInstanceInternalEndpointResponseStatus
}

func GetShowInstanceInternalEndpointResponseStatusEnum() ShowInstanceInternalEndpointResponseStatusEnum {
	return ShowInstanceInternalEndpointResponseStatusEnum{
		CREATING: ShowInstanceInternalEndpointResponseStatus{
			value: "Creating",
		},
		RUNNING: ShowInstanceInternalEndpointResponseStatus{
			value: "Running",
		},
		CREATE_ERROR: ShowInstanceInternalEndpointResponseStatus{
			value: "CreateError",
		},
		DELETING: ShowInstanceInternalEndpointResponseStatus{
			value: "Deleting",
		},
		DELETE_ERROR: ShowInstanceInternalEndpointResponseStatus{
			value: "DeleteError",
		},
	}
}

func (c ShowInstanceInternalEndpointResponseStatus) Value() string {
	return c.value
}

func (c ShowInstanceInternalEndpointResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceInternalEndpointResponseStatus) UnmarshalJSON(b []byte) error {
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
