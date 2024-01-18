package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobInstanceResponse Response Object
type ShowJobInstanceResponse struct {

	// 作业名称
	JobName *string `json:"jobName,omitempty"`

	// 作业实例ID
	InstanceId *int64 `json:"instanceId,omitempty"`

	// 作业实例状态： - waiting：等待运行 - running：运行中 - success：运行成功 - fail： 运行失败 - running-exception：运行异常 - pause： 暂停 - manual-stop：取消
	Status *string `json:"status,omitempty"`

	// 作业实例计划执行时间
	PlanTime *int64 `json:"planTime,omitempty"`

	// 作业实例实际执行开始时间
	StartTime *int64 `json:"startTime,omitempty"`

	// 作业实例实际执行结束时间
	EndTime *int64 `json:"endTime,omitempty"`

	// 执行耗时，单位：毫秒
	ExecuteTime *int64 `json:"executeTime,omitempty"`

	// 总的节点数据条数
	Total *int32 `json:"total,omitempty"`

	// 节点实例状态
	Nodes *[]NodeInstance `json:"nodes,omitempty"`

	// 作业调度方式： - 0：正常调度 - 2：手工调度 - 5：补数据 - 6：子作业调度 - 7：单次调度
	InstanceType *int32 `json:"instanceType,omitempty"`

	// 作业实例状态筛选为强制成功，默认值：false
	ForceSuccess *bool `json:"forceSuccess,omitempty"`

	// 作业实例状态筛选为忽略失败，默认值：false
	IgnoreSuccess  *bool `json:"ignoreSuccess,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowJobInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowJobInstanceResponse", string(data)}, " ")
}
