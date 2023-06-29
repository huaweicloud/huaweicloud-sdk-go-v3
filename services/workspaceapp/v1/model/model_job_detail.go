package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobDetail Job信息详情
type JobDetail struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// job类型,固定值1
	JobType *int32 `json:"job_type,omitempty"`

	// job执行状态 - 0：失败。（如果是开通失败，云运营查询到失败状态，直接退费给客户。如果是变更失败，当前是发运维工单，暂时还不是直接退费给客户） - 1：成功。（处理结果成功） - 2：处理中。 - 3：正在初始化。
	JobStatus *int32 `json:"job_status,omitempty"`

	// 子任务信息
	SubJobs *[]SubJobInfo `json:"sub_jobs,omitempty"`

	// Job处理开始时间
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// job处理结束时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 云服务预估的Job处理结束时间，只针对job有效，针对子job无效
	ExpectedEndTime *sdktime.SdkTime `json:"expected_end_time,omitempty"`

	// Job执行结果码
	ExecuteCode *string `json:"execute_code,omitempty"`

	// Job执行结果描述，以及每个SubJob的执行结果描述
	ExecuteMessage *string `json:"execute_message,omitempty"`
}

func (o JobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetail struct{}"
	}

	return strings.Join([]string{"JobDetail", string(data)}, " ")
}
