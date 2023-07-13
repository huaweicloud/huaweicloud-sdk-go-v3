package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateConnectionResponse Response Object
type UpdateConnectionResponse struct {

	// 事件源ID
	Id *string `json:"id,omitempty"`

	// 目标连接名称，租户下唯一，由小写字母、数字、点、下划线和中划线组成，必须以字母或数字开头，不能为default
	Name *string `json:"name,omitempty"`

	// 目标连接描述
	Description *string `json:"description,omitempty"`

	// 目标连接状态
	Status *UpdateConnectionResponseStatus `json:"status,omitempty"`

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

func (o UpdateConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateConnectionResponse", string(data)}, " ")
}

type UpdateConnectionResponseStatus struct {
	value string
}

type UpdateConnectionResponseStatusEnum struct {
	CREATING      UpdateConnectionResponseStatus
	CREATED       UpdateConnectionResponseStatus
	CREATE_FAILED UpdateConnectionResponseStatus
}

func GetUpdateConnectionResponseStatusEnum() UpdateConnectionResponseStatusEnum {
	return UpdateConnectionResponseStatusEnum{
		CREATING: UpdateConnectionResponseStatus{
			value: "CREATING",
		},
		CREATED: UpdateConnectionResponseStatus{
			value: "CREATED",
		},
		CREATE_FAILED: UpdateConnectionResponseStatus{
			value: "CREATE_FAILED",
		},
	}
}

func (c UpdateConnectionResponseStatus) Value() string {
	return c.value
}

func (c UpdateConnectionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateConnectionResponseStatus) UnmarshalJSON(b []byte) error {
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
