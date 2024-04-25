package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApisRequest Request Object
type ListApisRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ListApisRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// 是否返回专享版API的发布信息。
	XReturnPublishMessages *string `json:"x-return-publish-messages,omitempty"`

	// 查询起始坐标, 即跳过前X条数据。仅支持0或limit的整数倍，不满足则向下取整。
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据。
	Limit *int32 `json:"limit,omitempty"`

	// 根据API名称模糊查询。
	Name *string `json:"name,omitempty"`

	// 根据API描述信息模糊查询。
	Description *string `json:"description,omitempty"`

	// 根据API创建用户模糊查询。
	CreateUser *string `json:"create_user,omitempty"`

	// 根据API创建时间过滤，开始时间，如2024-02-24T16:00:00.000Z。
	StartTime *string `json:"start_time,omitempty"`

	// 根据API创建时间过滤，结束时间，如2024-04-05T15:59:59.998Z。
	EndTime *string `json:"end_time,omitempty"`

	// 标签。
	Tags *[]string `json:"tags,omitempty"`

	// API类型。
	ApiType *ListApisRequestApiType `json:"api_type,omitempty"`

	// API发布状态。
	PublishStatus *ListApisRequestPublishStatus `json:"publish_status,omitempty"`

	// 根据API用到的数据库表名模糊查询。
	TableName *string `json:"table_name,omitempty"`
}

func (o ListApisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisRequest struct{}"
	}

	return strings.Join([]string{"ListApisRequest", string(data)}, " ")
}

type ListApisRequestDlmType struct {
	value string
}

type ListApisRequestDlmTypeEnum struct {
	SHARED    ListApisRequestDlmType
	EXCLUSIVE ListApisRequestDlmType
}

func GetListApisRequestDlmTypeEnum() ListApisRequestDlmTypeEnum {
	return ListApisRequestDlmTypeEnum{
		SHARED: ListApisRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListApisRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListApisRequestDlmType) Value() string {
	return c.value
}

func (c ListApisRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApisRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ListApisRequestApiType struct {
	value string
}

type ListApisRequestApiTypeEnum struct {
	API_SPECIFIC_TYPE_CONFIGURATION ListApisRequestApiType
	API_SPECIFIC_TYPE_SCRIPT        ListApisRequestApiType
	API_SPECIFIC_TYPE_REGISTER      ListApisRequestApiType
	API_SPECIFIC_TYPE_MYBATIS       ListApisRequestApiType
	API_SPECIFIC_TYPE_GROOVY        ListApisRequestApiType
}

func GetListApisRequestApiTypeEnum() ListApisRequestApiTypeEnum {
	return ListApisRequestApiTypeEnum{
		API_SPECIFIC_TYPE_CONFIGURATION: ListApisRequestApiType{
			value: "API_SPECIFIC_TYPE_CONFIGURATION",
		},
		API_SPECIFIC_TYPE_SCRIPT: ListApisRequestApiType{
			value: "API_SPECIFIC_TYPE_SCRIPT",
		},
		API_SPECIFIC_TYPE_REGISTER: ListApisRequestApiType{
			value: "API_SPECIFIC_TYPE_REGISTER",
		},
		API_SPECIFIC_TYPE_MYBATIS: ListApisRequestApiType{
			value: "API_SPECIFIC_TYPE_MYBATIS",
		},
		API_SPECIFIC_TYPE_GROOVY: ListApisRequestApiType{
			value: "API_SPECIFIC_TYPE_GROOVY",
		},
	}
}

func (c ListApisRequestApiType) Value() string {
	return c.value
}

func (c ListApisRequestApiType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApisRequestApiType) UnmarshalJSON(b []byte) error {
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

type ListApisRequestPublishStatus struct {
	value string
}

type ListApisRequestPublishStatusEnum struct {
	PUBLISHED     ListApisRequestPublishStatus
	NOT_PUBLISHED ListApisRequestPublishStatus
}

func GetListApisRequestPublishStatusEnum() ListApisRequestPublishStatusEnum {
	return ListApisRequestPublishStatusEnum{
		PUBLISHED: ListApisRequestPublishStatus{
			value: "PUBLISHED",
		},
		NOT_PUBLISHED: ListApisRequestPublishStatus{
			value: "NOT_PUBLISHED",
		},
	}
}

func (c ListApisRequestPublishStatus) Value() string {
	return c.value
}

func (c ListApisRequestPublishStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApisRequestPublishStatus) UnmarshalJSON(b []byte) error {
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
