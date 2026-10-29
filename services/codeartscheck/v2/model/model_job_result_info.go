package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobResultInfo **参数解释：** 执行结果。
type JobResultInfo struct {

	// **参数解释：** 严重问题数。 **取值范围：** ≥0
	Critical *int32 `json:"CRITICAL,omitempty"`

	// **参数解释：** 延迟上线问题数。 **取值范围：** ≥0
	DelaystatusCount *int32 `json:"DELAYSTATUS_COUNT,omitempty"`

	// **参数解释：** 致命问题数。 **取值范围：** ≥0
	Fatal *int32 `json:"FATAL,omitempty"`

	// **参数解释：** 一般问题数。 **取值范围：** ≥0
	General *int32 `json:"GENERAL,omitempty"`

	// **参数解释：** 信息问题数。 **取值范围：** ≥0
	Info *int32 `json:"INFO,omitempty"`

	// **参数解释：** 提示问题数。 **取值范围：** ≥0
	Prompt *int32 `json:"PROMPT,omitempty"`

	// **参数解释：** 扫描代码行数，不包括空行和注释行。 **取值范围：** ≥0
	CodeLine *int32 `json:"codeLine,omitempty"`

	// **参数解释：** 缺陷总数。 **取值范围：** ≥0
	DefectCount *int32 `json:"defectCount,omitempty"`

	// **参数解释：** 已忽略问题数。 **取值范围：** ≥0
	Dismissed *int32 `json:"dismissed,omitempty"`

	// **参数解释：** 文件总数。 **取值范围：** ≥0
	FileTotal *int32 `json:"fileTotal,omitempty"`

	// **参数解释：** 已修复问题数。 **取值范围：** ≥0
	Fixed *int32 `json:"fixed,omitempty"`

	// **参数解释：** 执行ID，来自代码检查执行任务接口的返回值，每次执行代码检查任务都有唯一的执行ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 执行ID，来自代码检查执行任务接口的返回值，每次执行代码检查任务都有唯一的执行ID。 **取值范围：** 不涉及。
	JobId *string `json:"jobId,omitempty"`

	// **参数解释：** 未解决且新增的问题数。 **取值范围：** ≥0
	NewUnsolvedDefectCount *int32 `json:"newUnsolvedDefectCount,omitempty"`

	// **参数解释：** 非安全问题数。 **取值范围：** ≥0
	NonSecurityDefectCount *int32 `json:"nonSecurityDefectCount,omitempty"`

	// **参数解释：** 门禁结果。 **取值范围：** - success：门禁通过。 - error：门禁未通过。
	ReviewResult *string `json:"reviewResult,omitempty"`

	// **参数解释：** 安全问题数。 **取值范围：** ≥0
	SecurityDefectCount *int32 `json:"securityDefectCount,omitempty"`

	// **参数解释：** 已解决的问题数。 **取值范围：** 不涉及。
	SolvedDefectCount *int32 `json:"solvedDefectCount,omitempty"`

	// **参数解释：** 未解决的问题数。 **取值范围：** 不涉及。
	UnsolvedDefectCount *int32 `json:"unsolvedDefectCount,omitempty"`
}

func (o JobResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResultInfo struct{}"
	}

	return strings.Join([]string{"JobResultInfo", string(data)}, " ")
}
