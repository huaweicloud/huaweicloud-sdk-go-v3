package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DefectInfoV2 struct {

	// **参数解释：** MergeKey，问题ID。 **取值范围：** 不涉及。
	DefectId *string `json:"defect_id,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围：** 不涉及。
	DefectCheckerName *string `json:"defect_checker_name,omitempty"`

	// **参数解释：** 问题状态。 **取值范围：** - 0：待处理。 - 1：已解决。 - 2：已忽略。 [- 3：重新打开（需打开忽略审核开关）。](tag:fcs,hcs,hcs_sm,hcs_site,hcs_sitesm) [- 4：待审核（需打开忽略审核开关）。](tag:fcs,hcs,hcs_sm,hcs_site,hcs_sitesm)
	DefectStatus *string `json:"defect_status,omitempty"`

	// **参数解释：** 问题责任人。 **取值范围：** 不涉及。
	DefectOwner *string `json:"defect_owner,omitempty"`

	// **参数解释：** 标签，多个标签用逗号隔开。 **取值范围：** 不涉及。
	RuleSystemTags *string `json:"rule_system_tags,omitempty"`

	// **参数解释：** 规则ID。 **取值范围：** 不涉及。
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释：** 规则名称。 **取值范围：** 不涉及。
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释：** 问题所在文件行号。 **取值范围：** 不涉及。
	LineNumber *string `json:"line_number,omitempty"`

	// **参数解释：** 问题详细描述。 **取值范围：** 不涉及。
	DefectContent *string `json:"defect_content,omitempty"`

	// **参数解释：** 问题级别。 **取值范围：** - 0：致命。 - 1：严重。 - 2：一般。 - 3：提示。
	DefectLevel *string `json:"defect_level,omitempty"`

	// **参数解释：** 问题文件路径。 **取值范围：** 不涉及。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释：** 问题发现时间。使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围：** 不涉及。
	IssueKey *string `json:"issue_key,omitempty"`

	// **参数解释：** 代码片段内容。
	Fragment *[]DefectFragmentV2 `json:"fragment,omitempty"`

	// **参数解释：** 调用链信息。
	Events *[]DefectEvents `json:"events,omitempty"`
}

func (o DefectInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefectInfoV2 struct{}"
	}

	return strings.Join([]string{"DefectInfoV2", string(data)}, " ")
}
