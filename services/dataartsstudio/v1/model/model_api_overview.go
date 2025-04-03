package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiOverview struct {

	// API ID
	Id *string `json:"id,omitempty"`

	// API名称
	Name *string `json:"name,omitempty"`

	// API分组ID（共享版）
	GroupId *string `json:"group_id,omitempty"`

	// API描述
	Description *string `json:"description,omitempty"`

	// API状态（共享版）
	Status *ApiOverviewStatus `json:"status,omitempty"`

	// API调试状态（共享版）
	DebugStatus *ApiOverviewDebugStatus `json:"debug_status,omitempty"`

	// 发布信息列表（专享版）
	PublishMessages *[]ApiPublishDto `json:"publish_messages,omitempty"`

	// API 类型
	Type *ApiOverviewType `json:"type,omitempty"`

	// API审核人
	Manager *string `json:"manager,omitempty"`

	// API创建者
	CreateUser *string `json:"create_user,omitempty"`

	// API 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	AuthorizationStatus *ApiOverviewAuthorizationStatus `json:"authorization_status,omitempty"`
}

func (o ApiOverview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiOverview struct{}"
	}

	return strings.Join([]string{"ApiOverview", string(data)}, " ")
}

type ApiOverviewStatus struct {
	value string
}

type ApiOverviewStatusEnum struct {
	API_STATUS_CREATED             ApiOverviewStatus
	API_STATUS_PUBLISH_WAIT_REVIEW ApiOverviewStatus
	API_STATUS_PUBLISH_REJECT      ApiOverviewStatus
	API_STATUS_PUBLISHED           ApiOverviewStatus
	API_STATUS_WAITING_STOP        ApiOverviewStatus
	API_STATUS_STOPPED             ApiOverviewStatus
	API_STATUS_RECOVER_WAIT_REVIEW ApiOverviewStatus
	API_STATUS_WAITING_OFFLINE     ApiOverviewStatus
	API_STATUS_OFFLINE             ApiOverviewStatus
	API_STATUS_OFFLINE_WAIT_REVIEW ApiOverviewStatus
}

func GetApiOverviewStatusEnum() ApiOverviewStatusEnum {
	return ApiOverviewStatusEnum{
		API_STATUS_CREATED: ApiOverviewStatus{
			value: "API_STATUS_CREATED",
		},
		API_STATUS_PUBLISH_WAIT_REVIEW: ApiOverviewStatus{
			value: "API_STATUS_PUBLISH_WAIT_REVIEW",
		},
		API_STATUS_PUBLISH_REJECT: ApiOverviewStatus{
			value: "API_STATUS_PUBLISH_REJECT",
		},
		API_STATUS_PUBLISHED: ApiOverviewStatus{
			value: "API_STATUS_PUBLISHED",
		},
		API_STATUS_WAITING_STOP: ApiOverviewStatus{
			value: "API_STATUS_WAITING_STOP",
		},
		API_STATUS_STOPPED: ApiOverviewStatus{
			value: "API_STATUS_STOPPED",
		},
		API_STATUS_RECOVER_WAIT_REVIEW: ApiOverviewStatus{
			value: "API_STATUS_RECOVER_WAIT_REVIEW",
		},
		API_STATUS_WAITING_OFFLINE: ApiOverviewStatus{
			value: "API_STATUS_WAITING_OFFLINE",
		},
		API_STATUS_OFFLINE: ApiOverviewStatus{
			value: "API_STATUS_OFFLINE",
		},
		API_STATUS_OFFLINE_WAIT_REVIEW: ApiOverviewStatus{
			value: "API_STATUS_OFFLINE_WAIT_REVIEW",
		},
	}
}

func (c ApiOverviewStatus) Value() string {
	return c.value
}

func (c ApiOverviewStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiOverviewStatus) UnmarshalJSON(b []byte) error {
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

type ApiOverviewDebugStatus struct {
	value string
}

type ApiOverviewDebugStatusEnum struct {
	API_DEBUG_WAITING ApiOverviewDebugStatus
	API_DEBUG_FAILED  ApiOverviewDebugStatus
	API_DEBUG_SUCCESS ApiOverviewDebugStatus
}

func GetApiOverviewDebugStatusEnum() ApiOverviewDebugStatusEnum {
	return ApiOverviewDebugStatusEnum{
		API_DEBUG_WAITING: ApiOverviewDebugStatus{
			value: "API_DEBUG_WAITING",
		},
		API_DEBUG_FAILED: ApiOverviewDebugStatus{
			value: "API_DEBUG_FAILED",
		},
		API_DEBUG_SUCCESS: ApiOverviewDebugStatus{
			value: "API_DEBUG_SUCCESS",
		},
	}
}

func (c ApiOverviewDebugStatus) Value() string {
	return c.value
}

func (c ApiOverviewDebugStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiOverviewDebugStatus) UnmarshalJSON(b []byte) error {
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

type ApiOverviewType struct {
	value string
}

type ApiOverviewTypeEnum struct {
	API_SPECIFIC_TYPE_CONFIGURATION ApiOverviewType
	API_SPECIFIC_TYPE_SCRIPT        ApiOverviewType
	API_SPECIFIC_TYPE_REGISTER      ApiOverviewType
}

func GetApiOverviewTypeEnum() ApiOverviewTypeEnum {
	return ApiOverviewTypeEnum{
		API_SPECIFIC_TYPE_CONFIGURATION: ApiOverviewType{
			value: "API_SPECIFIC_TYPE_CONFIGURATION",
		},
		API_SPECIFIC_TYPE_SCRIPT: ApiOverviewType{
			value: "API_SPECIFIC_TYPE_SCRIPT",
		},
		API_SPECIFIC_TYPE_REGISTER: ApiOverviewType{
			value: "API_SPECIFIC_TYPE_REGISTER",
		},
	}
}

func (c ApiOverviewType) Value() string {
	return c.value
}

func (c ApiOverviewType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiOverviewType) UnmarshalJSON(b []byte) error {
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

type ApiOverviewAuthorizationStatus struct {
	value string
}

type ApiOverviewAuthorizationStatusEnum struct {
	NO_AUTHORIZATION_REQUIRED ApiOverviewAuthorizationStatus
	UNAUTHORIZED              ApiOverviewAuthorizationStatus
	AUTHORIZED                ApiOverviewAuthorizationStatus
}

func GetApiOverviewAuthorizationStatusEnum() ApiOverviewAuthorizationStatusEnum {
	return ApiOverviewAuthorizationStatusEnum{
		NO_AUTHORIZATION_REQUIRED: ApiOverviewAuthorizationStatus{
			value: "NO_AUTHORIZATION_REQUIRED",
		},
		UNAUTHORIZED: ApiOverviewAuthorizationStatus{
			value: "UNAUTHORIZED",
		},
		AUTHORIZED: ApiOverviewAuthorizationStatus{
			value: "AUTHORIZED",
		},
	}
}

func (c ApiOverviewAuthorizationStatus) Value() string {
	return c.value
}

func (c ApiOverviewAuthorizationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiOverviewAuthorizationStatus) UnmarshalJSON(b []byte) error {
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
