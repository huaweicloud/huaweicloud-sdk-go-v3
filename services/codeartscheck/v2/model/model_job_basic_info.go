package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobBasicInfo **参数解释：** 执行记录信息。
type JobBasicInfo struct {

	// **参数解释：** 执行人用户名。 **取值范围：** 不涉及。
	Executor *string `json:"executor,omitempty"`

	// **参数解释：** 执行记录的检查完成时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。 **取值范围：** 不涉及。
	FinishTime *string `json:"finishTime,omitempty"`

	JobResult *JobResultInfo `json:"jobResult,omitempty"`

	// **参数解释:** 任务类型。 **取值范围:** - FULL：版本级任务。 - INC_GIT：门禁级任务。 [- INC_COMMIT：流水线指定commitId增量检查、代码检查任务设置-检查范围-版本级检查模式的最近一次提交检查。](tag:hws,hws_hk) [- INC_COMMIT：代码检查任务设置-检查范围-版本级检查模式的最近一次提交检查。](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs) [- INC_LAST_SUCCESS_COMMIT：流水线指定commitId以及最后一次门禁成功的commitId间的增量检查。](tag:hws,hws_hk)
	JobType *string `json:"jobType,omitempty"`

	Process *ProcessInfo `json:"process,omitempty"`

	// **参数解释：**  执行记录的检查开始时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。 **取值范围：** 不涉及。
	StartTime *string `json:"startTime,omitempty"`

	// **参数解释:** 任务执行状态。 **取值范围:** - quering：初始状态。 - running：执行中。 - success：检查成功。 - failed：检查失败。 - aborted：任务中止。
	Status *string `json:"status,omitempty"`

	// **参数解释：**  任务ID，即新建检查任务接口的返回值，每个代码检查任务及其下的分支任务、通过API创建的增量任务都有唯一的任务ID。即[[查询任务列表](https://support.huaweicloud.com/api-codecheck/ShowTasksList.html)](tag:hws)[[查询任务列表](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ShowTasksList.html)](tag:hws_hk)[查询任务列表](tag:hcs,hcs_sm)接口响应中的id字段。 **取值范围：** 长度为[1~32]个字符。
	TaskId *string `json:"taskId,omitempty"`
}

func (o JobBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobBasicInfo struct{}"
	}

	return strings.Join([]string{"JobBasicInfo", string(data)}, " ")
}
