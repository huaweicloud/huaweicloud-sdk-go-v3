package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTasksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 每页显示条目数量，最大数量999，超过999后只返回999
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询，分页的偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int32 `json:"offset,omitempty"`

	// 任务ID，可为空
	TaskId *string `json:"task_id,omitempty"`

	// 模糊匹配任务名称，可为空
	Name *string `json:"name,omitempty"`

	// 任务状态，可为空 - 0 （停止/未启动） - 1 （运行中）
	Status *ListTasksRequestStatus `json:"status,omitempty"`

	// 任务类型 - realtime (实时) - timing (定时)
	TaskType *ListTasksRequestTaskType `json:"task_type,omitempty"`

	// 源端数据源ID，可为空
	SourceDatasourceId *string `json:"source_datasource_id,omitempty"`

	// 目标端数据源ID，可为空
	TargetDatasourceId *string `json:"target_datasource_id,omitempty"`

	// 查询排序的条件
	SortField *ListTasksRequestSortField `json:"sort_field,omitempty"`

	// 排序类型，可为空 - ASC (升序) - DESC (降序)
	SortType *ListTasksRequestSortType `json:"sort_type,omitempty"`

	// 执行状态，可为空 - UNSTARTED (未启动) - WAITING (等待执行) - RUNNING (执行中) - SUCCESS (执行成功) - CANCELLED (任务取消) - ERROR (执行异常)
	ExecuteStatus *string `json:"execute_status,omitempty"`

	// 源端数据源所属集成应用ID，可为空
	SourceAppId *string `json:"source_app_id,omitempty"`

	// 目标端数据源所属集成应用ID，可为空
	TargetAppId *string `json:"target_app_id,omitempty"`

	// 任务标签，可为空
	TaskTag *string `json:"task_tag,omitempty"`
}

func (o ListTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTasksRequest", string(data)}, " ")
}

type ListTasksRequestStatus struct {
	value int32
}

type ListTasksRequestStatusEnum struct {
	E_0 ListTasksRequestStatus
	E_1 ListTasksRequestStatus
}

func GetListTasksRequestStatusEnum() ListTasksRequestStatusEnum {
	return ListTasksRequestStatusEnum{
		E_0: ListTasksRequestStatus{
			value: 0,
		}, E_1: ListTasksRequestStatus{
			value: 1,
		},
	}
}

func (c ListTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListTasksRequestTaskType struct {
	value string
}

type ListTasksRequestTaskTypeEnum struct {
	REALTIME ListTasksRequestTaskType
	TIMING   ListTasksRequestTaskType
}

func GetListTasksRequestTaskTypeEnum() ListTasksRequestTaskTypeEnum {
	return ListTasksRequestTaskTypeEnum{
		REALTIME: ListTasksRequestTaskType{
			value: "realtime",
		},
		TIMING: ListTasksRequestTaskType{
			value: "timing",
		},
	}
}

func (c ListTasksRequestTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestTaskType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListTasksRequestSortField struct {
	value string
}

type ListTasksRequestSortFieldEnum struct {
	CREATED_DATE ListTasksRequestSortField
	STATUS       ListTasksRequestSortField
}

func GetListTasksRequestSortFieldEnum() ListTasksRequestSortFieldEnum {
	return ListTasksRequestSortFieldEnum{
		CREATED_DATE: ListTasksRequestSortField{
			value: "CREATED_DATE",
		},
		STATUS: ListTasksRequestSortField{
			value: "STATUS",
		},
	}
}

func (c ListTasksRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestSortField) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListTasksRequestSortType struct {
	value string
}

type ListTasksRequestSortTypeEnum struct {
	ASC  ListTasksRequestSortType
	DESC ListTasksRequestSortType
}

func GetListTasksRequestSortTypeEnum() ListTasksRequestSortTypeEnum {
	return ListTasksRequestSortTypeEnum{
		ASC: ListTasksRequestSortType{
			value: "ASC",
		},
		DESC: ListTasksRequestSortType{
			value: "DESC",
		},
	}
}

func (c ListTasksRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestSortType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
