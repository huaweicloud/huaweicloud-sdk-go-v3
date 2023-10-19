package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// AccessClientInfo 接入客户端信息
type AccessClientInfo struct {

	// 客户端ID
	Id *string `json:"id,omitempty"`

	// 客户端名称
	Name *string `json:"name,omitempty"`

	// 接入模式,SYSTEM,CUSTOM,AUTO
	AccessMode *AccessClientInfoAccessMode `json:"access_mode,omitempty"`

	// 客户端状态,CREATING,RUNNING,DELETING,DELETED,CREATE_FAIL,DELETE_FAIL
	Status *AccessClientInfoStatus `json:"status,omitempty"`

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 接入连接列表
	AccessConnections *[]AccessConnectionInfo `json:"access_connections,omitempty"`

	// 实例创建时间戳
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o AccessClientInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessClientInfo struct{}"
	}

	return strings.Join([]string{"AccessClientInfo", string(data)}, " ")
}

type AccessClientInfoAccessMode struct {
	value string
}

type AccessClientInfoAccessModeEnum struct {
	SYSTEM AccessClientInfoAccessMode
	CUSTOM AccessClientInfoAccessMode
	AUTO   AccessClientInfoAccessMode
}

func GetAccessClientInfoAccessModeEnum() AccessClientInfoAccessModeEnum {
	return AccessClientInfoAccessModeEnum{
		SYSTEM: AccessClientInfoAccessMode{
			value: "SYSTEM",
		},
		CUSTOM: AccessClientInfoAccessMode{
			value: "CUSTOM",
		},
		AUTO: AccessClientInfoAccessMode{
			value: "AUTO",
		},
	}
}

func (c AccessClientInfoAccessMode) Value() string {
	return c.value
}

func (c AccessClientInfoAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessClientInfoAccessMode) UnmarshalJSON(b []byte) error {
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

type AccessClientInfoStatus struct {
	value string
}

type AccessClientInfoStatusEnum struct {
	CREATING    AccessClientInfoStatus
	RUNNING     AccessClientInfoStatus
	DELETING    AccessClientInfoStatus
	DELETED     AccessClientInfoStatus
	CREATE_FAIL AccessClientInfoStatus
	DELETE_FAIL AccessClientInfoStatus
}

func GetAccessClientInfoStatusEnum() AccessClientInfoStatusEnum {
	return AccessClientInfoStatusEnum{
		CREATING: AccessClientInfoStatus{
			value: "CREATING",
		},
		RUNNING: AccessClientInfoStatus{
			value: "RUNNING",
		},
		DELETING: AccessClientInfoStatus{
			value: "DELETING",
		},
		DELETED: AccessClientInfoStatus{
			value: "DELETED",
		},
		CREATE_FAIL: AccessClientInfoStatus{
			value: "CREATE_FAIL",
		},
		DELETE_FAIL: AccessClientInfoStatus{
			value: "DELETE_FAIL",
		},
	}
}

func (c AccessClientInfoStatus) Value() string {
	return c.value
}

func (c AccessClientInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessClientInfoStatus) UnmarshalJSON(b []byte) error {
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
