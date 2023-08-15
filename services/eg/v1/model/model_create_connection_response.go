package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateConnectionResponse Response Object
type CreateConnectionResponse struct {

	// 事件源ID
	Id *string `json:"id,omitempty"`

	// 目标连接名称，租户下唯一，由小写字母、数字、点、下划线和中划线组成，必须以字母或数字开头，不能为default
	Name *string `json:"name,omitempty"`

	// 目标连接描述
	Description *string `json:"description,omitempty"`

	// 目标连接状态
	Status *CreateConnectionResponseStatus `json:"status,omitempty"`

	ErrorInfo *ErrorInfo `json:"error_info,omitempty"`

	// 待连接的VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 待连接的子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 私网目标连接使用的用户委托名称
	Agency *string `json:"agency,omitempty"`

	Flavor *ConnectionInfoFlavor `json:"flavor,omitempty"`

	Type *ConnectionType `json:"type,omitempty"`

	KafkaDetail *KafkaConnectionDetail `json:"kafka_detail,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectionResponse", string(data)}, " ")
}

type CreateConnectionResponseStatus struct {
	value string
}

type CreateConnectionResponseStatusEnum struct {
	CREATING      CreateConnectionResponseStatus
	CREATED       CreateConnectionResponseStatus
	CREATE_FAILED CreateConnectionResponseStatus
}

func GetCreateConnectionResponseStatusEnum() CreateConnectionResponseStatusEnum {
	return CreateConnectionResponseStatusEnum{
		CREATING: CreateConnectionResponseStatus{
			value: "CREATING",
		},
		CREATED: CreateConnectionResponseStatus{
			value: "CREATED",
		},
		CREATE_FAILED: CreateConnectionResponseStatus{
			value: "CREATE_FAILED",
		},
	}
}

func (c CreateConnectionResponseStatus) Value() string {
	return c.value
}

func (c CreateConnectionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConnectionResponseStatus) UnmarshalJSON(b []byte) error {
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
