package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineConcurrencyMgmt 流水线并发管理
type PipelineConcurrencyMgmt struct {

	// **参数解释**： 流水线ID，可以通过[查询流水线列表](ListPipelines.xml)接口，其中pipelines.pipelineId即为流水线ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符串。 **默认取值**： 不涉及。
	PipelineId *string `json:"pipeline_id,omitempty"`

	// **参数解释**： [流水线并发个数](tag:hws,hws_hk,hws_eu,ctc,cmcc,g42,sbc,hcs,hcs_site,hcs_sm,hcs_sitesm,fcs)[，最大并发受套餐和购买并发数限制](tag:hws,hws_hk,hws_eu,ctc,cmcc,g42,sbc)[。](tag:hws,hws_hk,hws_eu,ctc,cmcc,g42,sbc,hcs,hcs_site,hcs_sm,hcs_sitesm,fcs) **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConcurrencyNumber *int32 `json:"concurrency_number,omitempty"`

	// **参数解释**： 超出并发数时排队策略。 **约束限制**： 不涉及。 **取值范围**： - ABORT：忽略不执行。 - QUEUE：排队等待。 **默认取值**： 不涉及。
	ExceedAction *string `json:"exceed_action,omitempty"`

	// **参数解释**： 创建时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 更新时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**： 是否启用。 **约束限制**： 不涉及。 **取值范围**： - true：启用。 - false：未启用。 **默认取值**： 不涉及。
	Enable *bool `json:"enable,omitempty"`
}

func (o PipelineConcurrencyMgmt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineConcurrencyMgmt struct{}"
	}

	return strings.Join([]string{"PipelineConcurrencyMgmt", string(data)}, " ")
}
