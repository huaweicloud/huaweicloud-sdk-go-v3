package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunningRecord 训练作业运行及故障恢复记录。
type RunningRecord struct {

	// 本次运行开始时间的unix时间戳，单位为秒(s)。
	StartAt *int64 `json:"start_at,omitempty"`

	// 本次运行结束时间的unix时间戳，单位为秒(s)。
	EndAt *int64 `json:"end_at,omitempty"`

	// **参数解释**：本次运行加速卡启动时间的unix时间戳，单位为秒(s)。 **取值范围**：不涉及。
	XpuStartAt *int64 `json:"xpu_start_at,omitempty"`

	// 本地运行的启动方式： - init_or_rescheduled：代表本次启动为被调度后的首次运行，包括初次启动及调度恢复后的运行。 - restarted：代表本次启动非被调度后的首次运行，为进程重启后的运行。
	StartType *string `json:"start_type,omitempty"`

	// 本次运行结束原因。
	EndReason *string `json:"end_reason,omitempty"`

	// 引发本次运行结束的task worker ID(如worker-0)。
	EndRelatedTask *string `json:"end_related_task,omitempty"`

	// 本次运行结束后所采取的故障容忍策略，枚举值如下： - npu_proc_restart: NPU原地热恢复 - gpu_proc_restart: GPU原地热恢复 - proc_restart: 进程原地重启 - pod_reschedule: Pod级重调度 - job_reschedule: Job级重调度 - job_reschedule_with_taint: 隔离式Job重调度
	EndRecover *string `json:"end_recover,omitempty"`

	// 本次运行结束后在故障容忍策略降级前所采取的容忍策略，取值范围同end_recover。
	EndRecoverBeforeDowngrade *string `json:"end_recover_before_downgrade,omitempty"`

	// **参数解释**：本次运行异常结束时采取的所有故障容忍策略详情。
	RecoverRecords *[]RecoverRecord `json:"recover_records,omitempty"`
}

func (o RunningRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunningRecord struct{}"
	}

	return strings.Join([]string{"RunningRecord", string(data)}, " ")
}
