package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListAllAppRequestBody struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 分页页码， 表示从此页开始查询， page大于等于1
	Page int32 `json:"page"`

	// 每页显示的条目数量，size小于等于100
	Size int32 `json:"size"`

	// 排序字段的名称，当前仅支持name和startTime
	SortName *ListAllAppRequestBodySortName `json:"sort_name,omitempty"`

	// 排序顺序，正序（ASC）或者逆序（DESC)
	SortBy *string `json:"sort_by,omitempty"`

	// 应用状态列表，支持查询以下状态： abort: 部署中止 failed: 部署失败 not_started: 取消执行 pending: 排队中 running: 正在部署 succeeded: 部署成功 timeout: 部署超时 not_executed: 未执行
	States *[]ListAllAppRequestBodyStates `json:"states,omitempty"`

	// 应用的分组id，传入no_grouped为查询未分组的应用
	GroupId *string `json:"group_id,omitempty"`
}

func (o ListAllAppRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllAppRequestBody struct{}"
	}

	return strings.Join([]string{"ListAllAppRequestBody", string(data)}, " ")
}

type ListAllAppRequestBodySortName struct {
	value string
}

type ListAllAppRequestBodySortNameEnum struct {
	NAME       ListAllAppRequestBodySortName
	START_TIME ListAllAppRequestBodySortName
}

func GetListAllAppRequestBodySortNameEnum() ListAllAppRequestBodySortNameEnum {
	return ListAllAppRequestBodySortNameEnum{
		NAME: ListAllAppRequestBodySortName{
			value: "name",
		},
		START_TIME: ListAllAppRequestBodySortName{
			value: "startTime",
		},
	}
}

func (c ListAllAppRequestBodySortName) Value() string {
	return c.value
}

func (c ListAllAppRequestBodySortName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAllAppRequestBodySortName) UnmarshalJSON(b []byte) error {
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

type ListAllAppRequestBodyStates struct {
	value string
}

type ListAllAppRequestBodyStatesEnum struct {
	ABORT        ListAllAppRequestBodyStates
	FAILED       ListAllAppRequestBodyStates
	NOT_STARTED  ListAllAppRequestBodyStates
	PENDING      ListAllAppRequestBodyStates
	RUNNING      ListAllAppRequestBodyStates
	SUCCEEDED    ListAllAppRequestBodyStates
	TIMEOUT      ListAllAppRequestBodyStates
	NOT_EXECUTED ListAllAppRequestBodyStates
}

func GetListAllAppRequestBodyStatesEnum() ListAllAppRequestBodyStatesEnum {
	return ListAllAppRequestBodyStatesEnum{
		ABORT: ListAllAppRequestBodyStates{
			value: "abort",
		},
		FAILED: ListAllAppRequestBodyStates{
			value: "failed",
		},
		NOT_STARTED: ListAllAppRequestBodyStates{
			value: "not_started",
		},
		PENDING: ListAllAppRequestBodyStates{
			value: "pending",
		},
		RUNNING: ListAllAppRequestBodyStates{
			value: "running",
		},
		SUCCEEDED: ListAllAppRequestBodyStates{
			value: "succeeded",
		},
		TIMEOUT: ListAllAppRequestBodyStates{
			value: "timeout",
		},
		NOT_EXECUTED: ListAllAppRequestBodyStates{
			value: "not_executed",
		},
	}
}

func (c ListAllAppRequestBodyStates) Value() string {
	return c.value
}

func (c ListAllAppRequestBodyStates) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAllAppRequestBodyStates) UnmarshalJSON(b []byte) error {
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
