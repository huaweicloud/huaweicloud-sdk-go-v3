package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowJobDetailRequest Request Object
type ShowJobDetailRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowJobDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 任务详情类型。取值： - overview：任务概览信息。 - detail：任务基本信息。 - network：测试连接结果信息，需配合query_id参数一起查询。 - precheck：预检查结果信息，需配合query_id参数一起查询。 - progress：任务进度信息。 - log：任务日志信息，支持分页查询参数offset与limit。 - compare：查询对比任务。 - file：对象导入信息。 - is_writable：目标库解除只读结果。 - cloud_connection：录制回放他云连通性测试，需配合query_id参数一起查询。
	Type ShowJobDetailRequestType `json:"type"`

	// 通过指定Query ID查询任务详情。  说明：部分type类型的任务详情，需要通过触发该操作的请求返回的query_id进行操作结果查询。
	QueryId *string `json:"query_id,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。  说明：部分type类型的任务详情支持分页查询，可以通过传递该参数进行分页控制。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。  说明：部分type类型的任务详情支持分页查询，可以通过传递该参数进行分页控制。
	Limit *int32 `json:"limit,omitempty"`

	// 对比任务类型 - object_compare：对象对比。 - line_compare：行对比。 - content_compare：内容对比。 - data_compare：数据对比。
	CompareType *ShowJobDetailRequestCompareType `json:"compare_type,omitempty"`

	// 查询对比内容。取值： - overview：对比任务概览。 - list：数据对比任务列表。 - detail：对比详情。 - diff：不一致详情。
	QueryType *ShowJobDetailRequestQueryType `json:"query_type,omitempty"`

	// 查询对象对比详情类型。取值： - DB：库级对比详情。 - TABLE：表级对比详情。 - INDEX：索引对比详情。
	ObjectType *ShowJobDetailRequestObjectType `json:"object_type,omitempty"`

	// 对比任务ID。
	CompareTaskId *string `json:"compare_task_id,omitempty"`

	// 数据对比源库名称。
	SourceDbName *string `json:"source_db_name,omitempty"`

	// 数据对比目标库名称。
	TargetDbName *string `json:"target_db_name,omitempty"`

	// 对比结果类型。取值： - compare：对比完成。 - uncompare：无法对比。
	CompareDetailType *ShowJobDetailRequestCompareDetailType `json:"compare_detail_type,omitempty"`
}

func (o ShowJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowJobDetailRequest", string(data)}, " ")
}

type ShowJobDetailRequestXLanguage struct {
	value string
}

type ShowJobDetailRequestXLanguageEnum struct {
	EN_US ShowJobDetailRequestXLanguage
	ZH_CN ShowJobDetailRequestXLanguage
}

func GetShowJobDetailRequestXLanguageEnum() ShowJobDetailRequestXLanguageEnum {
	return ShowJobDetailRequestXLanguageEnum{
		EN_US: ShowJobDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowJobDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowJobDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowJobDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowJobDetailRequestType struct {
	value string
}

type ShowJobDetailRequestTypeEnum struct {
	OVERVIEW         ShowJobDetailRequestType
	DETAIL           ShowJobDetailRequestType
	NETWORK          ShowJobDetailRequestType
	PRECHECK         ShowJobDetailRequestType
	PROGRESS         ShowJobDetailRequestType
	LOG              ShowJobDetailRequestType
	COMPARE          ShowJobDetailRequestType
	FILE             ShowJobDetailRequestType
	IS_WRITABLE      ShowJobDetailRequestType
	CLOUD_CONNECTION ShowJobDetailRequestType
}

func GetShowJobDetailRequestTypeEnum() ShowJobDetailRequestTypeEnum {
	return ShowJobDetailRequestTypeEnum{
		OVERVIEW: ShowJobDetailRequestType{
			value: "overview",
		},
		DETAIL: ShowJobDetailRequestType{
			value: "detail",
		},
		NETWORK: ShowJobDetailRequestType{
			value: "network",
		},
		PRECHECK: ShowJobDetailRequestType{
			value: "precheck",
		},
		PROGRESS: ShowJobDetailRequestType{
			value: "progress",
		},
		LOG: ShowJobDetailRequestType{
			value: "log",
		},
		COMPARE: ShowJobDetailRequestType{
			value: "compare",
		},
		FILE: ShowJobDetailRequestType{
			value: "file",
		},
		IS_WRITABLE: ShowJobDetailRequestType{
			value: "is_writable",
		},
		CLOUD_CONNECTION: ShowJobDetailRequestType{
			value: "cloud_connection",
		},
	}
}

func (c ShowJobDetailRequestType) Value() string {
	return c.value
}

func (c ShowJobDetailRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobDetailRequestType) UnmarshalJSON(b []byte) error {
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

type ShowJobDetailRequestCompareType struct {
	value string
}

type ShowJobDetailRequestCompareTypeEnum struct {
	OBJECT_COMPARE  ShowJobDetailRequestCompareType
	LINE_COMPARE    ShowJobDetailRequestCompareType
	CONTENT_COMPARE ShowJobDetailRequestCompareType
	DATA_COMPARE    ShowJobDetailRequestCompareType
}

func GetShowJobDetailRequestCompareTypeEnum() ShowJobDetailRequestCompareTypeEnum {
	return ShowJobDetailRequestCompareTypeEnum{
		OBJECT_COMPARE: ShowJobDetailRequestCompareType{
			value: "object_compare",
		},
		LINE_COMPARE: ShowJobDetailRequestCompareType{
			value: "line_compare",
		},
		CONTENT_COMPARE: ShowJobDetailRequestCompareType{
			value: "content_compare",
		},
		DATA_COMPARE: ShowJobDetailRequestCompareType{
			value: "data_compare",
		},
	}
}

func (c ShowJobDetailRequestCompareType) Value() string {
	return c.value
}

func (c ShowJobDetailRequestCompareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobDetailRequestCompareType) UnmarshalJSON(b []byte) error {
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

type ShowJobDetailRequestQueryType struct {
	value string
}

type ShowJobDetailRequestQueryTypeEnum struct {
	OVERVIEW ShowJobDetailRequestQueryType
	LIST     ShowJobDetailRequestQueryType
	DETAIL   ShowJobDetailRequestQueryType
	DIFF     ShowJobDetailRequestQueryType
}

func GetShowJobDetailRequestQueryTypeEnum() ShowJobDetailRequestQueryTypeEnum {
	return ShowJobDetailRequestQueryTypeEnum{
		OVERVIEW: ShowJobDetailRequestQueryType{
			value: "overview",
		},
		LIST: ShowJobDetailRequestQueryType{
			value: "list",
		},
		DETAIL: ShowJobDetailRequestQueryType{
			value: "detail",
		},
		DIFF: ShowJobDetailRequestQueryType{
			value: "diff",
		},
	}
}

func (c ShowJobDetailRequestQueryType) Value() string {
	return c.value
}

func (c ShowJobDetailRequestQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobDetailRequestQueryType) UnmarshalJSON(b []byte) error {
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

type ShowJobDetailRequestObjectType struct {
	value string
}

type ShowJobDetailRequestObjectTypeEnum struct {
	DB    ShowJobDetailRequestObjectType
	TABLE ShowJobDetailRequestObjectType
	INDEX ShowJobDetailRequestObjectType
}

func GetShowJobDetailRequestObjectTypeEnum() ShowJobDetailRequestObjectTypeEnum {
	return ShowJobDetailRequestObjectTypeEnum{
		DB: ShowJobDetailRequestObjectType{
			value: "DB",
		},
		TABLE: ShowJobDetailRequestObjectType{
			value: "TABLE",
		},
		INDEX: ShowJobDetailRequestObjectType{
			value: "INDEX",
		},
	}
}

func (c ShowJobDetailRequestObjectType) Value() string {
	return c.value
}

func (c ShowJobDetailRequestObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobDetailRequestObjectType) UnmarshalJSON(b []byte) error {
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

type ShowJobDetailRequestCompareDetailType struct {
	value string
}

type ShowJobDetailRequestCompareDetailTypeEnum struct {
	COMPARE   ShowJobDetailRequestCompareDetailType
	UNCOMPARE ShowJobDetailRequestCompareDetailType
}

func GetShowJobDetailRequestCompareDetailTypeEnum() ShowJobDetailRequestCompareDetailTypeEnum {
	return ShowJobDetailRequestCompareDetailTypeEnum{
		COMPARE: ShowJobDetailRequestCompareDetailType{
			value: "compare",
		},
		UNCOMPARE: ShowJobDetailRequestCompareDetailType{
			value: "uncompare",
		},
	}
}

func (c ShowJobDetailRequestCompareDetailType) Value() string {
	return c.value
}

func (c ShowJobDetailRequestCompareDetailType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobDetailRequestCompareDetailType) UnmarshalJSON(b []byte) error {
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
