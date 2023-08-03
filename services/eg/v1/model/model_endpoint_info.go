package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EndpointInfo struct {

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
	Status *EndpointInfoStatus `json:"status,omitempty"`

	ErrorInfo *ErrorInfo `json:"error_info,omitempty"`

	// 访问端点类型
	Type *EndpointInfoType `json:"type,omitempty"`

	// 是否可更新
	Scalable *bool `json:"scalable,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 访问端点终端节点列表
	Endpoints *[]EndpointConnection `json:"endpoints,omitempty"`
}

func (o EndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointInfo struct{}"
	}

	return strings.Join([]string{"EndpointInfo", string(data)}, " ")
}

type EndpointInfoStatus struct {
	value string
}

type EndpointInfoStatusEnum struct {
	CREATED       EndpointInfoStatus
	CREATING      EndpointInfoStatus
	CREATE_FAILED EndpointInfoStatus
	DELETE_FAILED EndpointInfoStatus
}

func GetEndpointInfoStatusEnum() EndpointInfoStatusEnum {
	return EndpointInfoStatusEnum{
		CREATED: EndpointInfoStatus{
			value: "CREATED",
		},
		CREATING: EndpointInfoStatus{
			value: "CREATING",
		},
		CREATE_FAILED: EndpointInfoStatus{
			value: "CREATE_FAILED",
		},
		DELETE_FAILED: EndpointInfoStatus{
			value: "DELETE_FAILED",
		},
	}
}

func (c EndpointInfoStatus) Value() string {
	return c.value
}

func (c EndpointInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointInfoStatus) UnmarshalJSON(b []byte) error {
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

type EndpointInfoType struct {
	value string
}

type EndpointInfoTypeEnum struct {
	PRIVATE EndpointInfoType
	PUBLIC  EndpointInfoType
}

func GetEndpointInfoTypeEnum() EndpointInfoTypeEnum {
	return EndpointInfoTypeEnum{
		PRIVATE: EndpointInfoType{
			value: "PRIVATE",
		},
		PUBLIC: EndpointInfoType{
			value: "PUBLIC",
		},
	}
}

func (c EndpointInfoType) Value() string {
	return c.value
}

func (c EndpointInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointInfoType) UnmarshalJSON(b []byte) error {
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
