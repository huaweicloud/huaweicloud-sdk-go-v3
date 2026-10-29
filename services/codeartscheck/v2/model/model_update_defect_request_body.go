package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDefectRequestBody struct {

	// **参数解释：** MergeKey，问题ID。[[查询缺陷详情](ShowTaskDefectsV2.xml)](tag:hws,hws_hk)[查询缺陷详情](tag:fcs,hcs,hcs_sm,hcs_site,hcs_sitesm)响应中的defect_id字段 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DefectId *string `json:"defect_id,omitempty"`

	// **参数解释：** 问题状态。 **约束限制：** [不涉及。](tag:hws,hws_hk)[状态为0、1的问题无法直接修改到状态3。](tag:fcs,hcs,hcs_sm,hcs_site,hcs_sitesm) **取值范围：** - 0：待处理。 - 1：已解决。 - 2：已忽略。 [- 3：重新打开（需打开忽略审核开关）。](tag:fcs,hcs,hcs_sm,hcs_site,hcs_sitesm) [- 4：待审核（需打开忽略审核开关）。](tag:fcs,hcs,hcs_sm,hcs_site,hcs_sitesm) **默认取值：** 不涉及。
	DefectStatus *string `json:"defect_status,omitempty"`

	// **参数解释：** 问题忽略的审核人用户名。 **约束限制：** defect_status取值为4（需打开忽略审核开关）时，必填。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Approver *string `json:"approver,omitempty"`

	// **参数解释：** 本次修改问题状态操作的备注。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Comment *string `json:"comment,omitempty"`

	// **参数解释：** [忽略类型，将问题修改为待审核状态时选填。](tag:fcs,hcs,hcs_sm,hcs_site,hcs_sitesm)[废弃字段。](tag:hws,hws_hk) **约束限制：** 不涉及。 **取值范围：** - NO_RISK：评估无风险暂不解决。 - FALSE_POSITIVES：纯属工具误报。 - THIRD_PARTY：开源及第三方。 - EFFUCUENCY_RETAIN：性能敏感保留危险函数。 - NO_SUPPORT_INKERNEL：不支持的内核危险函数。 - EXCLUDE_SCENES：规范例外场景。 - COMPATIBLE_THIRD_PARTY：兼容第三方。 **默认取值：** NO_RISK
	IgnoreType *string `json:"ignoreType,omitempty"`
}

func (o UpdateDefectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefectRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDefectRequestBody", string(data)}, " ")
}
