package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecoverRecord struct {

	// **参数描述**：本次故障容忍策略开始执行时间的unix时间戳，单位为秒(s)，同时也是故障发生时间。 **取值范围**：不涉及。
	RecoverStartAt *int64 `json:"recover_start_at,omitempty"`

	// **参数描述**：本次故障容忍策略结束时间的unix时间戳，单位为秒(s)。 **取值范围**：不涉及。
	RecoverEndAt *int64 `json:"recover_end_at,omitempty"`

	// **参数描述**：本次故障容忍策略。 **取值范围**：枚举值如下： - npu_step_retry: Step重计算 - npu_proc_restart: NPU原地热恢复 - proc_restart: 进程原地重启 - pod_reschedule: Pod级重调度 - job_reschedule: Job级重调度 - job_reschedule_with_taint: 隔离式Job重调度
	Recover *string `json:"recover,omitempty"`

	// **参数描述**：本次故障场景。 **取值范围**：枚举值如下： - chip_fault: 芯片故障 - node_fault: 节点故障 - job_failed: 作业失败退出 - job_hanged: 作业卡死 - job_subhealth: 作业亚健康 - error_in_log: 日志异常
	FaultScenario *string `json:"fault_scenario,omitempty"`

	// **参数描述**：本次故障原因。 **取值范围**：不涉及。
	Reason *string `json:"reason,omitempty"`

	// **参数描述**：引发本次运行结束的task worker ID(如worker-0)。 **取值范围**：不涉及。
	RelatedTask *string `json:"related_task,omitempty"`

	// **参数描述**：本次故障执行结果。 **取值范围**：枚举值如下： - recovering: 执行中 - success: 成功 - failed: 失败 - downgrade: 策略降级 - terminated: 策略被终止 - quotaExceeded: 策略执行次数超限制
	RecoverResult *string `json:"recover_result,omitempty"`
}

func (o RecoverRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecoverRecord struct{}"
	}

	return strings.Join([]string{"RecoverRecord", string(data)}, " ")
}
