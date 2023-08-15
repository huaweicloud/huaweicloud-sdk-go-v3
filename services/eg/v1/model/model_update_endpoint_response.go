package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateEndpointResponse Response Object
type UpdateEndpointResponse struct {

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
	Status *UpdateEndpointResponseStatus `json:"status,omitempty"`

	ErrorInfo *ErrorInfo `json:"error_info,omitempty"`

	// 访问端点类型
	Type *UpdateEndpointResponseType `json:"type,omitempty"`

	// 是否可更新
	Scalable *bool `json:"scalable,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 访问端点终端节点列表
	Endpoints      *[]EndpointConnection `json:"endpoints,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UpdateEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointResponse", string(data)}, " ")
}

type UpdateEndpointResponseStatus struct {
	value string
}

type UpdateEndpointResponseStatusEnum struct {
	CREATED       UpdateEndpointResponseStatus
	CREATING      UpdateEndpointResponseStatus
	CREATE_FAILED UpdateEndpointResponseStatus
	DELETE_FAILED UpdateEndpointResponseStatus
}

func GetUpdateEndpointResponseStatusEnum() UpdateEndpointResponseStatusEnum {
	return UpdateEndpointResponseStatusEnum{
		CREATED: UpdateEndpointResponseStatus{
			value: "CREATED",
		},
		CREATING: UpdateEndpointResponseStatus{
			value: "CREATING",
		},
		CREATE_FAILED: UpdateEndpointResponseStatus{
			value: "CREATE_FAILED",
		},
		DELETE_FAILED: UpdateEndpointResponseStatus{
			value: "DELETE_FAILED",
		},
	}
}

func (c UpdateEndpointResponseStatus) Value() string {
	return c.value
}

func (c UpdateEndpointResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointResponseStatus) UnmarshalJSON(b []byte) error {
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

type UpdateEndpointResponseType struct {
	value string
}

type UpdateEndpointResponseTypeEnum struct {
	PRIVATE UpdateEndpointResponseType
	PUBLIC  UpdateEndpointResponseType
}

func GetUpdateEndpointResponseTypeEnum() UpdateEndpointResponseTypeEnum {
	return UpdateEndpointResponseTypeEnum{
		PRIVATE: UpdateEndpointResponseType{
			value: "PRIVATE",
		},
		PUBLIC: UpdateEndpointResponseType{
			value: "PUBLIC",
		},
	}
}

func (c UpdateEndpointResponseType) Value() string {
	return c.value
}

func (c UpdateEndpointResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointResponseType) UnmarshalJSON(b []byte) error {
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
