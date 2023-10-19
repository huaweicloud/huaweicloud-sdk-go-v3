package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowAccessClientResponse Response Object
type ShowAccessClientResponse struct {

	// 客户端ID
	Id *string `json:"id,omitempty"`

	// 客户端名称
	Name *string `json:"name,omitempty"`

	// 接入模式,SYSTEM,CUSTOM,AUTO
	AccessMode *ShowAccessClientResponseAccessMode `json:"access_mode,omitempty"`

	// 客户端状态,CREATING,RUNNING,DELETING,DELETED,CREATE_FAIL,DELETE_FAIL
	Status *ShowAccessClientResponseStatus `json:"status,omitempty"`

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

func (o ShowAccessClientResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessClientResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessClientResponse", string(data)}, " ")
}

type ShowAccessClientResponseAccessMode struct {
	value string
}

type ShowAccessClientResponseAccessModeEnum struct {
	SYSTEM ShowAccessClientResponseAccessMode
	CUSTOM ShowAccessClientResponseAccessMode
	AUTO   ShowAccessClientResponseAccessMode
}

func GetShowAccessClientResponseAccessModeEnum() ShowAccessClientResponseAccessModeEnum {
	return ShowAccessClientResponseAccessModeEnum{
		SYSTEM: ShowAccessClientResponseAccessMode{
			value: "SYSTEM",
		},
		CUSTOM: ShowAccessClientResponseAccessMode{
			value: "CUSTOM",
		},
		AUTO: ShowAccessClientResponseAccessMode{
			value: "AUTO",
		},
	}
}

func (c ShowAccessClientResponseAccessMode) Value() string {
	return c.value
}

func (c ShowAccessClientResponseAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAccessClientResponseAccessMode) UnmarshalJSON(b []byte) error {
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

type ShowAccessClientResponseStatus struct {
	value string
}

type ShowAccessClientResponseStatusEnum struct {
	CREATING    ShowAccessClientResponseStatus
	RUNNING     ShowAccessClientResponseStatus
	DELETING    ShowAccessClientResponseStatus
	DELETED     ShowAccessClientResponseStatus
	CREATE_FAIL ShowAccessClientResponseStatus
	DELETE_FAIL ShowAccessClientResponseStatus
}

func GetShowAccessClientResponseStatusEnum() ShowAccessClientResponseStatusEnum {
	return ShowAccessClientResponseStatusEnum{
		CREATING: ShowAccessClientResponseStatus{
			value: "CREATING",
		},
		RUNNING: ShowAccessClientResponseStatus{
			value: "RUNNING",
		},
		DELETING: ShowAccessClientResponseStatus{
			value: "DELETING",
		},
		DELETED: ShowAccessClientResponseStatus{
			value: "DELETED",
		},
		CREATE_FAIL: ShowAccessClientResponseStatus{
			value: "CREATE_FAIL",
		},
		DELETE_FAIL: ShowAccessClientResponseStatus{
			value: "DELETE_FAIL",
		},
	}
}

func (c ShowAccessClientResponseStatus) Value() string {
	return c.value
}

func (c ShowAccessClientResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAccessClientResponseStatus) UnmarshalJSON(b []byte) error {
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
