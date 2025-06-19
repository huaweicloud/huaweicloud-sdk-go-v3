package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSparkJobsRequest Request Object
type ListSparkJobsRequest struct {

	// 参数解释:   DLI队列名称，不填写则获取当前Project下所有批处理作业(不推荐使用) 示例: cluster1 约束限制:  匹配正则表达式'^(?!_)(?![0-9]+$)[A-Za-z0-9_]*$'的字符串 取值范围: 无 默认取值: 无
	ClusterName *string `json:"cluster_name,omitempty"`

	// 参数解释:   用于查询开始时间在该时间点之前的作业。时间格式为unix时间戳，单位：毫秒 示例: 156789546456 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	End *int64 `json:"end,omitempty"`

	// 参数解释:   起始批处理作业的索引号，默认从0开始 示例: 0 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	From *int32 `json:"from,omitempty"`

	// 参数解释:   批处理作业的名称 示例: dli_test 约束限制:  无 取值范围: 无 默认取值: 无
	JobName *string `json:"job_name,omitempty"`

	// 参数解释:   批处理作业的ID 示例: 03923a72-5ace-466a-a573-e8c7b08b8cf3 约束限制:  匹配正则表达式'^[A-Fa-f0-9_-]*$'的字符串 取值范围: 无 默认取值: 无
	JobId *string `json:"job-id,omitempty"`

	// 参数解释:   指定作业排序方式 示例: DURATION_DESC 约束限制:  无 取值范围: DURATION_DESC（作业运行时长降序） DURATION_ASC（作业运行时长升序） CREATE_TIME_DESC（作业提交时间降序） CREATE_TIME_ASC（作业提交时间升序） 默认取值: 无
	Order *string `json:"order,omitempty"`

	// 参数解释:   队列名称 示例: 03923a72-5ace-466a-a573-e8c7b08b8cf3 约束限制:  匹配正则表达式'^[A-Fa-f0-9_-]*$'的字符串 取值范围: 无 默认取值: 无
	QueueName *string `json:"queue_name,omitempty"`

	// 参数解释:   查询批处理作业的数量 示例: 1 约束限制:  无 取值范围: 无 默认取值: 100
	Size *int32 `json:"size,omitempty"`

	// 参数解释:   用于查询开始时间在该时间点之后的作业。时间格式为unix时间戳，单位：毫秒 示例: 156456784655 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	Start *int64 `json:"start,omitempty"`

	// 参数解释:   批处理作业的状态 示例: success 约束限制:  无 取值范围: starting（批处理作业正在启动） running（批处理作业正在执行任务） dead（批处理作业已退出） success（批处理作业执行成功） recovering（批处理作业正在恢复） 默认取值: 无
	State *ListSparkJobsRequestState `json:"state,omitempty"`
}

func (o ListSparkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkJobsRequest struct{}"
	}

	return strings.Join([]string{"ListSparkJobsRequest", string(data)}, " ")
}

type ListSparkJobsRequestState struct {
	value string
}

type ListSparkJobsRequestStateEnum struct {
	STARTING   ListSparkJobsRequestState
	RUNNING    ListSparkJobsRequestState
	DEAD       ListSparkJobsRequestState
	SUCCESS    ListSparkJobsRequestState
	RECOVERING ListSparkJobsRequestState
}

func GetListSparkJobsRequestStateEnum() ListSparkJobsRequestStateEnum {
	return ListSparkJobsRequestStateEnum{
		STARTING: ListSparkJobsRequestState{
			value: "starting",
		},
		RUNNING: ListSparkJobsRequestState{
			value: "running",
		},
		DEAD: ListSparkJobsRequestState{
			value: "dead",
		},
		SUCCESS: ListSparkJobsRequestState{
			value: "success",
		},
		RECOVERING: ListSparkJobsRequestState{
			value: "recovering",
		},
	}
}

func (c ListSparkJobsRequestState) Value() string {
	return c.value
}

func (c ListSparkJobsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSparkJobsRequestState) UnmarshalJSON(b []byte) error {
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
