package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConnectionInfo struct {

	// 事件源ID
	Id *string `json:"id,omitempty"`

	// 目标连接名称，租户下唯一，由小写字母、数字、点、下划线和中划线组成，必须以字母或数字开头，不能为default
	Name *string `json:"name,omitempty"`

	// 目标连接描述
	Description *string `json:"description,omitempty"`

	// 目标连接状态
	Status *ConnectionInfoStatus `json:"status,omitempty"`

	// 待连接的VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 待连接的子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 私网目标连接使用的用户委托名称
	Agency *string `json:"agency,omitempty"`

	Flavor *ConnectionInfoFlavor `json:"flavor,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`
}

func (o ConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionInfo struct{}"
	}

	return strings.Join([]string{"ConnectionInfo", string(data)}, " ")
}

type ConnectionInfoStatus struct {
	value string
}

type ConnectionInfoStatusEnum struct {
	CREATING      ConnectionInfoStatus
	CREATED       ConnectionInfoStatus
	CREATE_FAILED ConnectionInfoStatus
}

func GetConnectionInfoStatusEnum() ConnectionInfoStatusEnum {
	return ConnectionInfoStatusEnum{
		CREATING: ConnectionInfoStatus{
			value: "CREATING",
		},
		CREATED: ConnectionInfoStatus{
			value: "CREATED",
		},
		CREATE_FAILED: ConnectionInfoStatus{
			value: "CREATE_FAILED",
		},
	}
}

func (c ConnectionInfoStatus) Value() string {
	return c.value
}

func (c ConnectionInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionInfoStatus) UnmarshalJSON(b []byte) error {
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
