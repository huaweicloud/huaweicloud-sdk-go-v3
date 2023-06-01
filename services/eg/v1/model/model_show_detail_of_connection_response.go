package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowDetailOfConnectionResponse struct {

	// 事件源ID
	Id *string `json:"id,omitempty"`

	// 目标连接名称，租户下唯一，由小写字母、数字、点、下划线和中划线组成，必须以字母或数字开头，不能为default
	Name *string `json:"name,omitempty"`

	// 目标连接描述
	Description *string `json:"description,omitempty"`

	// 目标连接状态
	Status *ShowDetailOfConnectionResponseStatus `json:"status,omitempty"`

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
	UpdatedTime    *string `json:"updated_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailOfConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowDetailOfConnectionResponse", string(data)}, " ")
}

type ShowDetailOfConnectionResponseStatus struct {
	value string
}

type ShowDetailOfConnectionResponseStatusEnum struct {
	CREATING      ShowDetailOfConnectionResponseStatus
	CREATED       ShowDetailOfConnectionResponseStatus
	CREATE_FAILED ShowDetailOfConnectionResponseStatus
}

func GetShowDetailOfConnectionResponseStatusEnum() ShowDetailOfConnectionResponseStatusEnum {
	return ShowDetailOfConnectionResponseStatusEnum{
		CREATING: ShowDetailOfConnectionResponseStatus{
			value: "CREATING",
		},
		CREATED: ShowDetailOfConnectionResponseStatus{
			value: "CREATED",
		},
		CREATE_FAILED: ShowDetailOfConnectionResponseStatus{
			value: "CREATE_FAILED",
		},
	}
}

func (c ShowDetailOfConnectionResponseStatus) Value() string {
	return c.value
}

func (c ShowDetailOfConnectionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailOfConnectionResponseStatus) UnmarshalJSON(b []byte) error {
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
