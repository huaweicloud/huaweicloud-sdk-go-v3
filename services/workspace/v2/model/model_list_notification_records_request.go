package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNotificationRecordsRequest Request Object
type ListNotificationRecordsRequest struct {

	// 通过该类型查询桌面或桌面池相关的通知拦截记录 - DESKTOP: 查找当前projectId下与桌面相关的通知拦截记录 - DESKTOP_POOL: 查找当前projectId下与桌面池相关的通知拦截记录
	QueryType ListNotificationRecordsRequestQueryType `json:"query_type"`

	// 桌面名
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面池名称，桌面池名称必须保证唯一。桌面名称只允许输入中文、大写字母、小写字母、数字、中划线，长度范围为1~255。
	DesktopPoolName *string `json:"desktop_pool_name,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 过滤出与SMN通知类型一致的通知拦截记录 - EMAIL: 通过邮件查找通知拦截记录 - SMS: 通过短信查找通知拦截记录
	Type *ListNotificationRecordsRequestType `json:"type,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，取值范围0-1000，默认值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段名称，需要结合sort_type字段一起使用。 - operate_time 发送时间
	SortField *ListNotificationRecordsRequestSortField `json:"sort_field,omitempty"`

	// 排序类型，默认升序，需要结合sort_field字段一起使用。 - ASC 升序。 - DESC 降序。
	SortType *ListNotificationRecordsRequestSortType `json:"sort_type,omitempty"`
}

func (o ListNotificationRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationRecordsRequest", string(data)}, " ")
}

type ListNotificationRecordsRequestQueryType struct {
	value string
}

type ListNotificationRecordsRequestQueryTypeEnum struct {
	DESKTOP      ListNotificationRecordsRequestQueryType
	DESKTOP_POOL ListNotificationRecordsRequestQueryType
}

func GetListNotificationRecordsRequestQueryTypeEnum() ListNotificationRecordsRequestQueryTypeEnum {
	return ListNotificationRecordsRequestQueryTypeEnum{
		DESKTOP: ListNotificationRecordsRequestQueryType{
			value: "DESKTOP",
		},
		DESKTOP_POOL: ListNotificationRecordsRequestQueryType{
			value: "DESKTOP_POOL",
		},
	}
}

func (c ListNotificationRecordsRequestQueryType) Value() string {
	return c.value
}

func (c ListNotificationRecordsRequestQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationRecordsRequestQueryType) UnmarshalJSON(b []byte) error {
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

type ListNotificationRecordsRequestType struct {
	value string
}

type ListNotificationRecordsRequestTypeEnum struct {
	EMAIL ListNotificationRecordsRequestType
	SMS   ListNotificationRecordsRequestType
}

func GetListNotificationRecordsRequestTypeEnum() ListNotificationRecordsRequestTypeEnum {
	return ListNotificationRecordsRequestTypeEnum{
		EMAIL: ListNotificationRecordsRequestType{
			value: "EMAIL",
		},
		SMS: ListNotificationRecordsRequestType{
			value: "SMS",
		},
	}
}

func (c ListNotificationRecordsRequestType) Value() string {
	return c.value
}

func (c ListNotificationRecordsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationRecordsRequestType) UnmarshalJSON(b []byte) error {
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

type ListNotificationRecordsRequestSortField struct {
	value string
}

type ListNotificationRecordsRequestSortFieldEnum struct {
	OPERATE_TIME ListNotificationRecordsRequestSortField
}

func GetListNotificationRecordsRequestSortFieldEnum() ListNotificationRecordsRequestSortFieldEnum {
	return ListNotificationRecordsRequestSortFieldEnum{
		OPERATE_TIME: ListNotificationRecordsRequestSortField{
			value: "operate_time",
		},
	}
}

func (c ListNotificationRecordsRequestSortField) Value() string {
	return c.value
}

func (c ListNotificationRecordsRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationRecordsRequestSortField) UnmarshalJSON(b []byte) error {
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

type ListNotificationRecordsRequestSortType struct {
	value string
}

type ListNotificationRecordsRequestSortTypeEnum struct {
	ASC  ListNotificationRecordsRequestSortType
	DESC ListNotificationRecordsRequestSortType
}

func GetListNotificationRecordsRequestSortTypeEnum() ListNotificationRecordsRequestSortTypeEnum {
	return ListNotificationRecordsRequestSortTypeEnum{
		ASC: ListNotificationRecordsRequestSortType{
			value: "ASC",
		},
		DESC: ListNotificationRecordsRequestSortType{
			value: "DESC",
		},
	}
}

func (c ListNotificationRecordsRequestSortType) Value() string {
	return c.value
}

func (c ListNotificationRecordsRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationRecordsRequestSortType) UnmarshalJSON(b []byte) error {
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
