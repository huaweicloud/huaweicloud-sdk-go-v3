package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApiResponse Response Object
type ShowApiResponse struct {

	// API的ID
	Id *string `json:"id,omitempty"`

	// API名称
	Name *string `json:"name,omitempty"`

	// API所属分组的ID（共享版）
	GroupId *string `json:"group_id,omitempty"`

	// API 描述
	Description *string `json:"description,omitempty"`

	// API 访问协议
	Protocol *ShowApiResponseProtocol `json:"protocol,omitempty"`

	// 发布类型，公开或者私有
	PublishType *ShowApiResponsePublishType `json:"publish_type,omitempty"`

	// 是否开启日志记录
	LogFlag *bool `json:"log_flag,omitempty"`

	// API的访问路径
	Path *string `json:"path,omitempty"`

	// 共享版域名
	Host *string `json:"host,omitempty"`

	Hosts *InstanceHostDto `json:"hosts,omitempty"`

	// API访问方式
	RequestType *ShowApiResponseRequestType `json:"request_type,omitempty"`

	// API创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// API 审核人名称
	Manager *string `json:"manager,omitempty"`

	// API的状态（共享版）
	Status *ShowApiResponseStatus `json:"status,omitempty"`

	// API 类型
	Type *ShowApiResponseType `json:"type,omitempty"`

	// API调试状态（共享版）
	DebugStatus *ShowApiResponseDebugStatus `json:"debug_status,omitempty"`

	// 发布信息列表（专享版）
	PublishMessages *[]ApiPublishDto `json:"publish_messages,omitempty"`

	// API请求参数
	RequestParas *[]RequestPara `json:"request_paras,omitempty"`

	DatasourceConfig *DatasourceConfig `json:"datasource_config,omitempty"`

	BackendConfig  *BackendConfig `json:"backend_config,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiResponse struct{}"
	}

	return strings.Join([]string{"ShowApiResponse", string(data)}, " ")
}

type ShowApiResponseProtocol struct {
	value string
}

type ShowApiResponseProtocolEnum struct {
	PROTOCOL_TYPE_HTTP  ShowApiResponseProtocol
	PROTOCOL_TYPE_HTTPS ShowApiResponseProtocol
}

func GetShowApiResponseProtocolEnum() ShowApiResponseProtocolEnum {
	return ShowApiResponseProtocolEnum{
		PROTOCOL_TYPE_HTTP: ShowApiResponseProtocol{
			value: "PROTOCOL_TYPE_HTTP",
		},
		PROTOCOL_TYPE_HTTPS: ShowApiResponseProtocol{
			value: "PROTOCOL_TYPE_HTTPS",
		},
	}
}

func (c ShowApiResponseProtocol) Value() string {
	return c.value
}

func (c ShowApiResponseProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiResponseProtocol) UnmarshalJSON(b []byte) error {
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

type ShowApiResponsePublishType struct {
	value string
}

type ShowApiResponsePublishTypeEnum struct {
	PUBLISH_TYPE_PUBLIC  ShowApiResponsePublishType
	PUBLISH_TYPE_PRIVATE ShowApiResponsePublishType
}

func GetShowApiResponsePublishTypeEnum() ShowApiResponsePublishTypeEnum {
	return ShowApiResponsePublishTypeEnum{
		PUBLISH_TYPE_PUBLIC: ShowApiResponsePublishType{
			value: "PUBLISH_TYPE_PUBLIC",
		},
		PUBLISH_TYPE_PRIVATE: ShowApiResponsePublishType{
			value: "PUBLISH_TYPE_PRIVATE",
		},
	}
}

func (c ShowApiResponsePublishType) Value() string {
	return c.value
}

func (c ShowApiResponsePublishType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiResponsePublishType) UnmarshalJSON(b []byte) error {
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

type ShowApiResponseRequestType struct {
	value string
}

type ShowApiResponseRequestTypeEnum struct {
	REQUEST_TYPE_POST ShowApiResponseRequestType
	REQUEST_TYPE_GET  ShowApiResponseRequestType
}

func GetShowApiResponseRequestTypeEnum() ShowApiResponseRequestTypeEnum {
	return ShowApiResponseRequestTypeEnum{
		REQUEST_TYPE_POST: ShowApiResponseRequestType{
			value: "REQUEST_TYPE_POST",
		},
		REQUEST_TYPE_GET: ShowApiResponseRequestType{
			value: "REQUEST_TYPE_GET",
		},
	}
}

func (c ShowApiResponseRequestType) Value() string {
	return c.value
}

func (c ShowApiResponseRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiResponseRequestType) UnmarshalJSON(b []byte) error {
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

type ShowApiResponseStatus struct {
	value string
}

type ShowApiResponseStatusEnum struct {
	API_STATUS_CREATED             ShowApiResponseStatus
	API_STATUS_PUBLISH_WAIT_REVIEW ShowApiResponseStatus
	API_STATUS_PUBLISH_REJECT      ShowApiResponseStatus
	API_STATUS_PUBLISHED           ShowApiResponseStatus
	API_STATUS_WAITING_STOP        ShowApiResponseStatus
	API_STATUS_STOPPED             ShowApiResponseStatus
	API_STATUS_RECOVER_WAIT_REVIEW ShowApiResponseStatus
	API_STATUS_WAITING_OFFLINE     ShowApiResponseStatus
	API_STATUS_OFFLINE             ShowApiResponseStatus
	API_STATUS_OFFLINE_WAIT_REVIEW ShowApiResponseStatus
}

func GetShowApiResponseStatusEnum() ShowApiResponseStatusEnum {
	return ShowApiResponseStatusEnum{
		API_STATUS_CREATED: ShowApiResponseStatus{
			value: "API_STATUS_CREATED",
		},
		API_STATUS_PUBLISH_WAIT_REVIEW: ShowApiResponseStatus{
			value: "API_STATUS_PUBLISH_WAIT_REVIEW",
		},
		API_STATUS_PUBLISH_REJECT: ShowApiResponseStatus{
			value: "API_STATUS_PUBLISH_REJECT",
		},
		API_STATUS_PUBLISHED: ShowApiResponseStatus{
			value: "API_STATUS_PUBLISHED",
		},
		API_STATUS_WAITING_STOP: ShowApiResponseStatus{
			value: "API_STATUS_WAITING_STOP",
		},
		API_STATUS_STOPPED: ShowApiResponseStatus{
			value: "API_STATUS_STOPPED",
		},
		API_STATUS_RECOVER_WAIT_REVIEW: ShowApiResponseStatus{
			value: "API_STATUS_RECOVER_WAIT_REVIEW",
		},
		API_STATUS_WAITING_OFFLINE: ShowApiResponseStatus{
			value: "API_STATUS_WAITING_OFFLINE",
		},
		API_STATUS_OFFLINE: ShowApiResponseStatus{
			value: "API_STATUS_OFFLINE",
		},
		API_STATUS_OFFLINE_WAIT_REVIEW: ShowApiResponseStatus{
			value: "API_STATUS_OFFLINE_WAIT_REVIEW",
		},
	}
}

func (c ShowApiResponseStatus) Value() string {
	return c.value
}

func (c ShowApiResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowApiResponseType struct {
	value string
}

type ShowApiResponseTypeEnum struct {
	API_SPECIFIC_TYPE_CONFIGURATION ShowApiResponseType
	API_SPECIFIC_TYPE_SCRIPT        ShowApiResponseType
	API_SPECIFIC_TYPE_REGISTER      ShowApiResponseType
}

func GetShowApiResponseTypeEnum() ShowApiResponseTypeEnum {
	return ShowApiResponseTypeEnum{
		API_SPECIFIC_TYPE_CONFIGURATION: ShowApiResponseType{
			value: "API_SPECIFIC_TYPE_CONFIGURATION",
		},
		API_SPECIFIC_TYPE_SCRIPT: ShowApiResponseType{
			value: "API_SPECIFIC_TYPE_SCRIPT",
		},
		API_SPECIFIC_TYPE_REGISTER: ShowApiResponseType{
			value: "API_SPECIFIC_TYPE_REGISTER",
		},
	}
}

func (c ShowApiResponseType) Value() string {
	return c.value
}

func (c ShowApiResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiResponseType) UnmarshalJSON(b []byte) error {
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

type ShowApiResponseDebugStatus struct {
	value string
}

type ShowApiResponseDebugStatusEnum struct {
	API_DEBUG_WAITING ShowApiResponseDebugStatus
	API_DEBUG_FAILED  ShowApiResponseDebugStatus
	API_DEBUG_SUCCESS ShowApiResponseDebugStatus
}

func GetShowApiResponseDebugStatusEnum() ShowApiResponseDebugStatusEnum {
	return ShowApiResponseDebugStatusEnum{
		API_DEBUG_WAITING: ShowApiResponseDebugStatus{
			value: "API_DEBUG_WAITING",
		},
		API_DEBUG_FAILED: ShowApiResponseDebugStatus{
			value: "API_DEBUG_FAILED",
		},
		API_DEBUG_SUCCESS: ShowApiResponseDebugStatus{
			value: "API_DEBUG_SUCCESS",
		},
	}
}

func (c ShowApiResponseDebugStatus) Value() string {
	return c.value
}

func (c ShowApiResponseDebugStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiResponseDebugStatus) UnmarshalJSON(b []byte) error {
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
