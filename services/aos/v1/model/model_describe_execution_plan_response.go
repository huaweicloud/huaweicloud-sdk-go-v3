package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type DescribeExecutionPlanResponse struct {

	// 栈的唯一Id,为uuid
	StackId *string `json:"stack_id,omitempty"`

	// 栈的名字
	StackName *string `json:"stack_name,omitempty"`

	// 执行计划的唯一Id，由资源编排服务随机生成,为uuid
	ExecutionPlanId *string `json:"execution_plan_id,omitempty"`

	// 执行计划的名字
	ExecutionPlanName *string `json:"execution_plan_name,omitempty"`

	// 执行计划的描述，此描述为用户在创建执行计划时指定
	Description *string `json:"description,omitempty"`

	// 参数列表
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`

	// 创建该执行计划时，给出的vars文件中的内容
	VarsUriContent *string `json:"vars_uri_content,omitempty"`

	// 创建该执行计划时，请求Body体中的tfvars文件内容
	VarsBody *string `json:"vars_body,omitempty"`

	// 执行计划的生成时间
	CreateTime *string `json:"create_time,omitempty"`

	// 执行时间
	ApplyTime *string `json:"apply_time,omitempty"`

	// 执行计划的执行状态，只有当AVAILABLE的时候才可以使用apply执行     * `CREATION_IN_PROGRESS` - 正在生成     * `CREATION_FAILED` - 生成失败     * `AVAILABLE` - 执行计划已经生成完成。可以使用apply进行执行     * `APPLY_IN_PROGRESS` - 执行计划正在执行     * `APPLIED` - 执行完成
	Status *DescribeExecutionPlanResponseStatus `json:"status,omitempty"`

	// 展示执行计划状态更多细节的信息
	StatusMessage *string `json:"status_message,omitempty"`

	Summary        *ExecutionPlanSummary `json:"summary,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o DescribeExecutionPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeExecutionPlanResponse struct{}"
	}

	return strings.Join([]string{"DescribeExecutionPlanResponse", string(data)}, " ")
}

type DescribeExecutionPlanResponseStatus struct {
	value string
}

type DescribeExecutionPlanResponseStatusEnum struct {
	CREATION_IN_PROGRESS DescribeExecutionPlanResponseStatus
	CREATION_FAILED      DescribeExecutionPlanResponseStatus
	AVAILABLE            DescribeExecutionPlanResponseStatus
	APPLY_IN_PROGRESS    DescribeExecutionPlanResponseStatus
	APPLIED              DescribeExecutionPlanResponseStatus
}

func GetDescribeExecutionPlanResponseStatusEnum() DescribeExecutionPlanResponseStatusEnum {
	return DescribeExecutionPlanResponseStatusEnum{
		CREATION_IN_PROGRESS: DescribeExecutionPlanResponseStatus{
			value: "CREATION_IN_PROGRESS",
		},
		CREATION_FAILED: DescribeExecutionPlanResponseStatus{
			value: "CREATION_FAILED",
		},
		AVAILABLE: DescribeExecutionPlanResponseStatus{
			value: "AVAILABLE",
		},
		APPLY_IN_PROGRESS: DescribeExecutionPlanResponseStatus{
			value: "APPLY_IN_PROGRESS",
		},
		APPLIED: DescribeExecutionPlanResponseStatus{
			value: "APPLIED",
		},
	}
}

func (c DescribeExecutionPlanResponseStatus) Value() string {
	return c.value
}

func (c DescribeExecutionPlanResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DescribeExecutionPlanResponseStatus) UnmarshalJSON(b []byte) error {
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
