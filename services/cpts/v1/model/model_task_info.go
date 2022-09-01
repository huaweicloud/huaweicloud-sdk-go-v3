package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInfo struct {

	// bench_concurrent
	BenchConcurrent *int32 `json:"bench_concurrent,omitempty" xml:"bench_concurrent"`

	// case_list
	CaseList *[]CaseInfo `json:"case_list,omitempty" xml:"case_list"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`

	// description
	Description *string `json:"description,omitempty" xml:"description"`

	// name
	Name *string `json:"name,omitempty" xml:"name"`

	// operate_mode
	OperateMode *int32 `json:"operate_mode,omitempty" xml:"operate_mode"`

	// project_id
	ProjectId *int32 `json:"project_id,omitempty" xml:"project_id"`

	// related_temp_running_data
	RelatedTempRunningData *[]RelatedTempRunningData `json:"related_temp_running_data,omitempty" xml:"related_temp_running_data"`

	// run_status
	RunStatus *int32 `json:"run_status,omitempty" xml:"run_status"`

	// update_time
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// parallel
	Parallel *bool `json:"parallel,omitempty" xml:"parallel"`
}

func (o TaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInfo struct{}"
	}

	return strings.Join([]string{"TaskInfo", string(data)}, " ")
}
