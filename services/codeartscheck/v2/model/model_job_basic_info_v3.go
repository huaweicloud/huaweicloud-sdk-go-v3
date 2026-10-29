package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobBasicInfoV3 **参数解释：** 执行记录信息。
type JobBasicInfoV3 struct {

	// **参数解释：** 执行人用户名。 **取值范围：** 长度为[1~255]个字符。
	Executor *string `json:"executor,omitempty"`

	// **参数解释：** 执行记录的检查完成时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。 **取值范围：** 不涉及。
	FinishTime *string `json:"finishTime,omitempty"`

	// **参数解释：** 执行ID，来自代码检查执行任务接口的返回值，每次执行代码检查任务都有唯一的执行ID。 **取值范围：** 长度为[1~32]个字符。
	Id *string `json:"id,omitempty"`

	JobResult *JobResultInfo `json:"jobResult,omitempty"`

	// **参数解释:** 任务类型。 **取值范围:** - FULL：版本级任务。 - INC_GIT：门禁级任务。 [- INC_COMMIT：流水线指定commitId增量检查、代码检查任务设置-检查范围-版本级检查模式的最近一次提交检查。](tag:hws,hws_hk) [- INC_COMMIT：代码检查任务设置-检查范围-版本级检查模式的最近一次提交检查。](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs) [- INC_LAST_SUCCESS_COMMIT：流水线指定commitId以及最后一次门禁成功的commitId间的增量检查。](tag:hws,hws_hk)
	JobType *string `json:"jobType,omitempty"`

	Process *ProcessInfo `json:"process,omitempty"`

	// **参数解释：** 门禁结果。 **取值范围：** - success：门禁通过。 - error：门禁未通过。
	ReviewResult *string `json:"reviewResult,omitempty"`

	// **参数解释：** 规则集配置列表。
	RuleSets *[]ConfigTemplateRuleSetInfo `json:"ruleSets,omitempty"`

	// **参数解释：** 执行来源。 **取值范围：** - customer：来自页面点击。 - Scheduled：来自定时任务。 - CodeCheckEdge：来自Edge服务。 - CodeArtsPipeline：来自流水线。 - DevCloudWebhook：来自CodeArtsRepo的Webhook。 - GitCodeActionPipeline：来自流水线的GitCodeAction插件。
	Source *string `json:"source,omitempty"`

	// **参数解释：** 源分支的代码仓地址。 **取值范围：** 长度为[1~1024]个字符。
	SourceRepoUrl *string `json:"sourceRepoUrl,omitempty"`

	// **参数解释：**  执行记录的检查开始时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。 **取值范围：** 不涉及。
	StartTime *string `json:"startTime,omitempty"`

	// **参数解释:** 任务执行状态。 **取值范围:** - quering：初始状态。 - running：执行中。 - success：检查成功。 - failed：检查失败。 - aborted：任务中止。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 子任务ID，该字段一般返回空数组[]。
	SubJobIds *[]string `json:"subJobIds,omitempty"`

	// **参数解释：** 子任务状态，该字段一般返回空数组[]。
	SubJobStatuses *[]string `json:"subJobStatuses,omitempty"`

	// **参数解释：** 目标分支的代码仓地址。 **取值范围：** 长度为[1~1024]个字符。
	TargetRepoUrl *string `json:"targetRepoUrl,omitempty"`

	// **参数解释：** 任务ID，即新建检查任务接口的返回值，每个代码检查任务及其下的分支任务、通过API创建的增量任务都有唯一的任务ID。即[[查询任务列表](https://support.huaweicloud.com/api-codecheck/ShowTasksList.html)](tag:hws)[[查询任务列表](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ShowTasksList.html)](tag:hws_hk)[查询任务列表](tag:hcs,hcs_sm)接口响应中的id字段。 **取值范围：** 长度为[1~32]个字符。
	TaskId *string `json:"taskId,omitempty"`

	// **参数解释：** 任务总执行时长，格式为%dm%ds，例如\"6m6s\"。 **取值范围：** 不涉及。
	TotalExecuteTime *string `json:"totalExecuteTime,omitempty"`
}

func (o JobBasicInfoV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobBasicInfoV3 struct{}"
	}

	return strings.Join([]string{"JobBasicInfoV3", string(data)}, " ")
}
