package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEnvironmentsRequest Request Object
type ListEnvironmentsRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 分页页码， 表示从此页开始查询， page大于等于1
	PageIndex *int32 `json:"page_index,omitempty"`

	// 每页显示的条目数量，size小于等于100
	PageSize *int32 `json:"page_size,omitempty"`

	// 要查询的环境名称
	Name *string `json:"name,omitempty"`

	// 排序字段，支持按照环境名称|用户名称|创建时间|用户昵称排序
	SortKey *ListEnvironmentsRequestSortKey `json:"sort_key,omitempty"`

	// 排序顺序，DESC降序，ASC升序
	SortDir *ListEnvironmentsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListEnvironmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentsRequest struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsRequest", string(data)}, " ")
}

type ListEnvironmentsRequestSortKey struct {
	value string
}

type ListEnvironmentsRequestSortKeyEnum struct {
	NAME         ListEnvironmentsRequestSortKey
	USER_NAME    ListEnvironmentsRequestSortKey
	CREATED_TIME ListEnvironmentsRequestSortKey
	NICK_NAME    ListEnvironmentsRequestSortKey
}

func GetListEnvironmentsRequestSortKeyEnum() ListEnvironmentsRequestSortKeyEnum {
	return ListEnvironmentsRequestSortKeyEnum{
		NAME: ListEnvironmentsRequestSortKey{
			value: "NAME",
		},
		USER_NAME: ListEnvironmentsRequestSortKey{
			value: "USER_NAME",
		},
		CREATED_TIME: ListEnvironmentsRequestSortKey{
			value: "CREATED_TIME",
		},
		NICK_NAME: ListEnvironmentsRequestSortKey{
			value: "NICK_NAME",
		},
	}
}

func (c ListEnvironmentsRequestSortKey) Value() string {
	return c.value
}

func (c ListEnvironmentsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnvironmentsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListEnvironmentsRequestSortDir struct {
	value string
}

type ListEnvironmentsRequestSortDirEnum struct {
	DESC ListEnvironmentsRequestSortDir
	ASC  ListEnvironmentsRequestSortDir
}

func GetListEnvironmentsRequestSortDirEnum() ListEnvironmentsRequestSortDirEnum {
	return ListEnvironmentsRequestSortDirEnum{
		DESC: ListEnvironmentsRequestSortDir{
			value: "DESC",
		},
		ASC: ListEnvironmentsRequestSortDir{
			value: "ASC",
		},
	}
}

func (c ListEnvironmentsRequestSortDir) Value() string {
	return c.value
}

func (c ListEnvironmentsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnvironmentsRequestSortDir) UnmarshalJSON(b []byte) error {
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
