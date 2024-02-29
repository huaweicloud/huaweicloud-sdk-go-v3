package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFactoryJobInstancesByNameRequest Request Object
type ListFactoryJobInstancesByNameRequest struct {

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`

	// 有Body体的情况下必须，无Body体的情况下则无需填写和校验，默认值：application/json
	ContentType *string `json:"Content-Type,omitempty"`

	// 分页返回结果，指定每页最大记录数。 范围[1,1000] 默认值：10
	Limit *int32 `json:"limit,omitempty"`

	// 分页的起始页，默认值为0。取值范围大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 返回作业实例开始时间大于min_plain_time的作业实例，单位为毫秒ms，默认设置为查询当天0点，最大可支持查询一个月。
	MinPlanTime *int64 `json:"min_plan_time,omitempty"`

	// 返回作业实例开始时间小于max_plain_time的作业实例，单位为毫秒ms，默认设置为当前时间。
	MaxPlanTime *int64 `json:"max_plan_time,omitempty"`

	// 实例运行状态: - waiting：等待运行 - running：运行中 - success：运行成功 - fail： 运行失败 - running-exception：运行异常 - pause： 暂停 - manual-stop：取消 - skip-by-depend：跳过 - freeze：冻结 默认查全部
	Status *ListFactoryJobInstancesByNameRequestStatus `json:"status,omitempty"`

	// 作业名称
	JobName string `json:"job_name"`

	// status为success的时候使用，true则筛选出强制成功的作业实例默认值：false
	ForceSuccess *bool `json:"force_success,omitempty"`

	// status为success的时候使用，true则筛选出忽略失败的作业实例默认值：false
	IgnoreSuccess *bool `json:"ignore_success,omitempty"`

	// 作业调度方式: -0:正常调度 -2:手工调度 -5:补数据 -6:子作业调度 -7:单次调度
	InstanceType *string `json:"instance_type,omitempty"`
}

func (o ListFactoryJobInstancesByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryJobInstancesByNameRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryJobInstancesByNameRequest", string(data)}, " ")
}

type ListFactoryJobInstancesByNameRequestStatus struct {
	value string
}

type ListFactoryJobInstancesByNameRequestStatusEnum struct {
	WAITING           ListFactoryJobInstancesByNameRequestStatus
	RUNNING           ListFactoryJobInstancesByNameRequestStatus
	SUCCESS           ListFactoryJobInstancesByNameRequestStatus
	FAIL              ListFactoryJobInstancesByNameRequestStatus
	RUNNING_EXCEPTION ListFactoryJobInstancesByNameRequestStatus
	PAUSE             ListFactoryJobInstancesByNameRequestStatus
	MANUAL_STOP       ListFactoryJobInstancesByNameRequestStatus
	SKIP_BY_DEPEND    ListFactoryJobInstancesByNameRequestStatus
	FREEZE            ListFactoryJobInstancesByNameRequestStatus
}

func GetListFactoryJobInstancesByNameRequestStatusEnum() ListFactoryJobInstancesByNameRequestStatusEnum {
	return ListFactoryJobInstancesByNameRequestStatusEnum{
		WAITING: ListFactoryJobInstancesByNameRequestStatus{
			value: "waiting",
		},
		RUNNING: ListFactoryJobInstancesByNameRequestStatus{
			value: "running",
		},
		SUCCESS: ListFactoryJobInstancesByNameRequestStatus{
			value: "success",
		},
		FAIL: ListFactoryJobInstancesByNameRequestStatus{
			value: "fail",
		},
		RUNNING_EXCEPTION: ListFactoryJobInstancesByNameRequestStatus{
			value: "running-exception",
		},
		PAUSE: ListFactoryJobInstancesByNameRequestStatus{
			value: "pause",
		},
		MANUAL_STOP: ListFactoryJobInstancesByNameRequestStatus{
			value: "manual-stop",
		},
		SKIP_BY_DEPEND: ListFactoryJobInstancesByNameRequestStatus{
			value: "skip-by-depend",
		},
		FREEZE: ListFactoryJobInstancesByNameRequestStatus{
			value: "freeze",
		},
	}
}

func (c ListFactoryJobInstancesByNameRequestStatus) Value() string {
	return c.value
}

func (c ListFactoryJobInstancesByNameRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFactoryJobInstancesByNameRequestStatus) UnmarshalJSON(b []byte) error {
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
