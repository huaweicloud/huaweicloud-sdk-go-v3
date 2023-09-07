package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListJobInstancesRequest Request Object
type ListJobInstancesRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 分页参数:每页限定数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数：页数
	Offset *int32 `json:"offset,omitempty"`

	// 返回作业实际开始时间大于minPlanTime的作业实例，单位为毫秒ms。
	MinPlanTime *int64 `json:"minPlanTime,omitempty"`

	// 返回作业实际开始时间小于maxPlanTime的作业实例，单位为毫秒ms。
	MaxPlanTime *int64 `json:"maxPlanTime,omitempty"`

	// 实例运行状态： - waiting：等待运行 - running：运行中 - success：运行成功 - fail： 运行失败 - running-exception：运行异常 - pause： 暂停 - manual-stop：取消
	Status *ListJobInstancesRequestStatus `json:"status,omitempty"`

	// 支持通过作业名进行精确查询。
	PreciseQuery *bool `json:"preciseQuery,omitempty"`

	// 作业名称。如果要查询指定批处理作业的实例列表，jobName就是批处理作业名称；如果要查询实时作业下某个节点关联的子作业，jobName格式为[实时作业名称]_[节点名称]。
	JobName *string `json:"jobName,omitempty"`

	// 作业调度方式： - 0：正常调度 - 2：手工调度 - 5：补数据 - 6：子作业调度 - 7：单次调度
	InstanceType *ListJobInstancesRequestInstanceType `json:"instanceType,omitempty"`
}

func (o ListJobInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListJobInstancesRequest", string(data)}, " ")
}

type ListJobInstancesRequestStatus struct {
	value string
}

type ListJobInstancesRequestStatusEnum struct {
	WAITING           ListJobInstancesRequestStatus
	RUNNING           ListJobInstancesRequestStatus
	SUCCESS           ListJobInstancesRequestStatus
	FAIL              ListJobInstancesRequestStatus
	RUNNING_EXCEPTION ListJobInstancesRequestStatus
	PAUSE             ListJobInstancesRequestStatus
	MANUAL_STOP       ListJobInstancesRequestStatus
}

func GetListJobInstancesRequestStatusEnum() ListJobInstancesRequestStatusEnum {
	return ListJobInstancesRequestStatusEnum{
		WAITING: ListJobInstancesRequestStatus{
			value: "waiting",
		},
		RUNNING: ListJobInstancesRequestStatus{
			value: "running",
		},
		SUCCESS: ListJobInstancesRequestStatus{
			value: "success",
		},
		FAIL: ListJobInstancesRequestStatus{
			value: "fail",
		},
		RUNNING_EXCEPTION: ListJobInstancesRequestStatus{
			value: "running-exception",
		},
		PAUSE: ListJobInstancesRequestStatus{
			value: "pause",
		},
		MANUAL_STOP: ListJobInstancesRequestStatus{
			value: "manual-stop",
		},
	}
}

func (c ListJobInstancesRequestStatus) Value() string {
	return c.value
}

func (c ListJobInstancesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobInstancesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListJobInstancesRequestInstanceType struct {
	value string
}

type ListJobInstancesRequestInstanceTypeEnum struct {
	E_0 ListJobInstancesRequestInstanceType
	E_2 ListJobInstancesRequestInstanceType
	E_5 ListJobInstancesRequestInstanceType
	E_6 ListJobInstancesRequestInstanceType
	E_7 ListJobInstancesRequestInstanceType
}

func GetListJobInstancesRequestInstanceTypeEnum() ListJobInstancesRequestInstanceTypeEnum {
	return ListJobInstancesRequestInstanceTypeEnum{
		E_0: ListJobInstancesRequestInstanceType{
			value: "0",
		},
		E_2: ListJobInstancesRequestInstanceType{
			value: "2",
		},
		E_5: ListJobInstancesRequestInstanceType{
			value: "5",
		},
		E_6: ListJobInstancesRequestInstanceType{
			value: "6",
		},
		E_7: ListJobInstancesRequestInstanceType{
			value: "7",
		},
	}
}

func (c ListJobInstancesRequestInstanceType) Value() string {
	return c.value
}

func (c ListJobInstancesRequestInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobInstancesRequestInstanceType) UnmarshalJSON(b []byte) error {
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
