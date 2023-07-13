package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPublishDto struct {

	// 发布编号
	Id *string `json:"id,omitempty"`

	// api编号
	ApiId *string `json:"api_id,omitempty"`

	// 集群编号
	InstanceId *string `json:"instance_id,omitempty"`

	// 集群名称
	InstanceName *string `json:"instance_name,omitempty"`

	// api状态
	ApiStatus *ApiPublishDtoApiStatus `json:"api_status,omitempty"`

	// api调试状态
	ApiDebug *ApiPublishDtoApiDebug `json:"api_debug,omitempty"`
}

func (o ApiPublishDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPublishDto struct{}"
	}

	return strings.Join([]string{"ApiPublishDto", string(data)}, " ")
}

type ApiPublishDtoApiStatus struct {
	value string
}

type ApiPublishDtoApiStatusEnum struct {
	API_STATUS_CREATED             ApiPublishDtoApiStatus
	API_STATUS_PUBLISH_WAIT_REVIEW ApiPublishDtoApiStatus
	API_STATUS_PUBLISH_REJECT      ApiPublishDtoApiStatus
	API_STATUS_PUBLISHED           ApiPublishDtoApiStatus
	API_STATUS_WAITING_STOP        ApiPublishDtoApiStatus
	API_STATUS_STOPPED             ApiPublishDtoApiStatus
	API_STATUS_RECOVER_WAIT_REVIEW ApiPublishDtoApiStatus
	API_STATUS_WAITING_OFFLINE     ApiPublishDtoApiStatus
	API_STATUS_OFFLINE             ApiPublishDtoApiStatus
}

func GetApiPublishDtoApiStatusEnum() ApiPublishDtoApiStatusEnum {
	return ApiPublishDtoApiStatusEnum{
		API_STATUS_CREATED: ApiPublishDtoApiStatus{
			value: "API_STATUS_CREATED",
		},
		API_STATUS_PUBLISH_WAIT_REVIEW: ApiPublishDtoApiStatus{
			value: "API_STATUS_PUBLISH_WAIT_REVIEW",
		},
		API_STATUS_PUBLISH_REJECT: ApiPublishDtoApiStatus{
			value: "API_STATUS_PUBLISH_REJECT",
		},
		API_STATUS_PUBLISHED: ApiPublishDtoApiStatus{
			value: "API_STATUS_PUBLISHED",
		},
		API_STATUS_WAITING_STOP: ApiPublishDtoApiStatus{
			value: "API_STATUS_WAITING_STOP",
		},
		API_STATUS_STOPPED: ApiPublishDtoApiStatus{
			value: "API_STATUS_STOPPED",
		},
		API_STATUS_RECOVER_WAIT_REVIEW: ApiPublishDtoApiStatus{
			value: "API_STATUS_RECOVER_WAIT_REVIEW",
		},
		API_STATUS_WAITING_OFFLINE: ApiPublishDtoApiStatus{
			value: "API_STATUS_WAITING_OFFLINE",
		},
		API_STATUS_OFFLINE: ApiPublishDtoApiStatus{
			value: "API_STATUS_OFFLINE",
		},
	}
}

func (c ApiPublishDtoApiStatus) Value() string {
	return c.value
}

func (c ApiPublishDtoApiStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPublishDtoApiStatus) UnmarshalJSON(b []byte) error {
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

type ApiPublishDtoApiDebug struct {
	value string
}

type ApiPublishDtoApiDebugEnum struct {
	API_DEBUG_WAITING ApiPublishDtoApiDebug
	API_DEBUG_FAILED  ApiPublishDtoApiDebug
	API_DEBUG_SUCCESS ApiPublishDtoApiDebug
}

func GetApiPublishDtoApiDebugEnum() ApiPublishDtoApiDebugEnum {
	return ApiPublishDtoApiDebugEnum{
		API_DEBUG_WAITING: ApiPublishDtoApiDebug{
			value: "API_DEBUG_WAITING",
		},
		API_DEBUG_FAILED: ApiPublishDtoApiDebug{
			value: "API_DEBUG_FAILED",
		},
		API_DEBUG_SUCCESS: ApiPublishDtoApiDebug{
			value: "API_DEBUG_SUCCESS",
		},
	}
}

func (c ApiPublishDtoApiDebug) Value() string {
	return c.value
}

func (c ApiPublishDtoApiDebug) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPublishDtoApiDebug) UnmarshalJSON(b []byte) error {
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
