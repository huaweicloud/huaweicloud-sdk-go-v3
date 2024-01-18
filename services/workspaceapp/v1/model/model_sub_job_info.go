package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubJobInfo 子任务信息。
type SubJobInfo struct {

	// 子job标识。
	JobId *string `json:"job_id,omitempty"`

	// 任务类型，固定值2：子Job。
	JobType *int32 `json:"job_type,omitempty"`

	// job状态 - 0：失败。 - 1：成功。 - 2：处理中。 - 3：正在初始化。
	JobStatus *int32 `json:"job_status,omitempty"`

	// SubJob中处理的云服务/云资源对象。 创建、规格变更等场景是必填；冻结、解冻、删除等场景可空。
	Entities *[]JobResourceInfo `json:"entities,omitempty"`

	// 任务开始时间。
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 任务结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 云服务预估的Job处理结束时间，只针对job有效，针对子job无效。
	ExpectedEndTime *sdktime.SdkTime `json:"expected_end_time,omitempty"`

	// Job执行结果码。
	ExecuteCode *string `json:"execute_code,omitempty"`

	// Job执行结果描述，以及每个SubJob的执行结果描述。
	ExecuteMessage *string `json:"execute_message,omitempty"`
}

func (o SubJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobInfo struct{}"
	}

	return strings.Join([]string{"SubJobInfo", string(data)}, " ")
}
