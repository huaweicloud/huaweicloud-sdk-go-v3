package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleTaskInfoV2 struct {

	// **参数解释：** 任务下的默认分支任务的ID。 **取值范围：** 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释：** 任务名称。 **取值范围：** 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释：** 任务创建者的用户ID。 **取值范围：** 不涉及。
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释：** 代码仓地址。 **取值范围：** 不涉及。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释：** 任务的默认分支，非代码仓默认分支。 **取值范围：** 不涉及。
	GitBranch *string `json:"git_branch,omitempty"`

	// **参数解释：** 任务创建时间。 **取值范围：** UTC格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 任务下所有分支任务、增量任务的最近检查时间。 **取值范围：** UTC格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如\"2025-11-28T06:55:41Z\"。
	LastCheckTime *string `json:"last_check_time,omitempty"`

	// **参数解释：** [密级ID,可见信息密级的ID，从页面接口GET https://{endpoint}/projects/v1/options的响应参数中获取，其中的code值即为密级ID。](tag:hcs_sm,hcs_sitesm)[该参数暂无业务意义，忽略。](tag:fcs) **取值范围：** [不涉及,取值根据用户配置的密级情况变化。](tag:hcs_sm,hcs_sitesm)[不涉及。](tag:fcs)
	SmLevel *string `json:"sm_level,omitempty"`

	// **参数解释：** [密级英文名称，取值根据用户配置的密级情况变化。](tag:hcs_sm,hcs_sitesm)[该参数暂无业务意义，忽略。](tag:fcs) **取值范围：** 不涉及。
	SmLevelNameEn *string `json:"sm_level_name_en,omitempty"`

	// **参数解释：** [密级中文名称，取值根据用户配置的密级情况变化。](tag:hcs_sm,hcs_sitesm)[该参数暂无业务意义，忽略。](tag:fcs) **取值范围：** 不涉及。
	SmLevelName *string `json:"sm_level_name,omitempty"`
}

func (o SimpleTaskInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleTaskInfoV2 struct{}"
	}

	return strings.Join([]string{"SimpleTaskInfoV2", string(data)}, " ")
}
