package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateEndpointResponse struct {

	// 访问端点ID
	Id *string `json:"id,omitempty"`

	// 用户指定的访问端点名称
	Name *string `json:"name,omitempty"`

	// 访问端点所在的VPC的ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 访问端点所在的子网的ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 访问域名
	Domain *string `json:"domain,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 访问端点状态
	Status *CreateEndpointResponseStatus `json:"status,omitempty"`

	// 访问端点类型
	Type *CreateEndpointResponseType `json:"type,omitempty"`

	// 是否可更新
	Scalable *bool `json:"scalable,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 访问端点终端节点列表
	Endpoints *[]EndpointConnection `json:"endpoints,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateEndpointResponse", string(data)}, " ")
}

type CreateEndpointResponseStatus struct {
	value string
}

type CreateEndpointResponseStatusEnum struct {
	CREATED       CreateEndpointResponseStatus
	CREATING      CreateEndpointResponseStatus
	CREATE_FAILED CreateEndpointResponseStatus
	DELETE_FAILED CreateEndpointResponseStatus
}

func GetCreateEndpointResponseStatusEnum() CreateEndpointResponseStatusEnum {
	return CreateEndpointResponseStatusEnum{
		CREATED: CreateEndpointResponseStatus{
			value: "CREATED",
		},
		CREATING: CreateEndpointResponseStatus{
			value: "CREATING",
		},
		CREATE_FAILED: CreateEndpointResponseStatus{
			value: "CREATE_FAILED",
		},
		DELETE_FAILED: CreateEndpointResponseStatus{
			value: "DELETE_FAILED",
		},
	}
}

func (c CreateEndpointResponseStatus) Value() string {
	return c.value
}

func (c CreateEndpointResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateEndpointResponseType struct {
	value string
}

type CreateEndpointResponseTypeEnum struct {
	PRIVATE CreateEndpointResponseType
	PUBLIC  CreateEndpointResponseType
}

func GetCreateEndpointResponseTypeEnum() CreateEndpointResponseTypeEnum {
	return CreateEndpointResponseTypeEnum{
		PRIVATE: CreateEndpointResponseType{
			value: "PRIVATE",
		},
		PUBLIC: CreateEndpointResponseType{
			value: "PUBLIC",
		},
	}
}

func (c CreateEndpointResponseType) Value() string {
	return c.value
}

func (c CreateEndpointResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointResponseType) UnmarshalJSON(b []byte) error {
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
