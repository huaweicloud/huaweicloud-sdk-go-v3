package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateAccessClientResponse Response Object
type CreateAccessClientResponse struct {

	// 客户端ID
	Id *string `json:"id,omitempty"`

	// 客户端名称
	Name *string `json:"name,omitempty"`

	// 接入模式,SYSTEM,CUSTOM,AUTO
	AccessMode *CreateAccessClientResponseAccessMode `json:"access_mode,omitempty"`

	// 客户端状态,CREATING,RUNNING,DELETING,DELETED,CREATE_FAIL,DELETE_FAIL
	Status *CreateAccessClientResponseStatus `json:"status,omitempty"`

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 接入连接列表
	AccessConnections *[]AccessConnectionInfo `json:"access_connections,omitempty"`

	// 实例创建时间戳
	CreateTime     *sdktime.SdkTime `json:"create_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateAccessClientResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessClientResponse struct{}"
	}

	return strings.Join([]string{"CreateAccessClientResponse", string(data)}, " ")
}

type CreateAccessClientResponseAccessMode struct {
	value string
}

type CreateAccessClientResponseAccessModeEnum struct {
	SYSTEM CreateAccessClientResponseAccessMode
	CUSTOM CreateAccessClientResponseAccessMode
	AUTO   CreateAccessClientResponseAccessMode
}

func GetCreateAccessClientResponseAccessModeEnum() CreateAccessClientResponseAccessModeEnum {
	return CreateAccessClientResponseAccessModeEnum{
		SYSTEM: CreateAccessClientResponseAccessMode{
			value: "SYSTEM",
		},
		CUSTOM: CreateAccessClientResponseAccessMode{
			value: "CUSTOM",
		},
		AUTO: CreateAccessClientResponseAccessMode{
			value: "AUTO",
		},
	}
}

func (c CreateAccessClientResponseAccessMode) Value() string {
	return c.value
}

func (c CreateAccessClientResponseAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAccessClientResponseAccessMode) UnmarshalJSON(b []byte) error {
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

type CreateAccessClientResponseStatus struct {
	value string
}

type CreateAccessClientResponseStatusEnum struct {
	CREATING    CreateAccessClientResponseStatus
	RUNNING     CreateAccessClientResponseStatus
	DELETING    CreateAccessClientResponseStatus
	DELETED     CreateAccessClientResponseStatus
	CREATE_FAIL CreateAccessClientResponseStatus
	DELETE_FAIL CreateAccessClientResponseStatus
}

func GetCreateAccessClientResponseStatusEnum() CreateAccessClientResponseStatusEnum {
	return CreateAccessClientResponseStatusEnum{
		CREATING: CreateAccessClientResponseStatus{
			value: "CREATING",
		},
		RUNNING: CreateAccessClientResponseStatus{
			value: "RUNNING",
		},
		DELETING: CreateAccessClientResponseStatus{
			value: "DELETING",
		},
		DELETED: CreateAccessClientResponseStatus{
			value: "DELETED",
		},
		CREATE_FAIL: CreateAccessClientResponseStatus{
			value: "CREATE_FAIL",
		},
		DELETE_FAIL: CreateAccessClientResponseStatus{
			value: "DELETE_FAIL",
		},
	}
}

func (c CreateAccessClientResponseStatus) Value() string {
	return c.value
}

func (c CreateAccessClientResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAccessClientResponseStatus) UnmarshalJSON(b []byte) error {
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
