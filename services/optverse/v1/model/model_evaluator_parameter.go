package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EvaluatorParameter struct {

	// **参数解释**： 使用的llm模型列表。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,30]。 **默认取值**： 不涉及
	LlmModels []string `json:"llm_models"`

	// **参数解释**： 最大评估数量。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,32]。 **默认取值**： 不涉及
	EvaluatorMaxWorkers int32 `json:"evaluator_max_workers"`

	// **参数解释**： 最大生成数量。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,32]。 **默认取值**： 不涉及
	SearchMaxWorkers int32 `json:"search_max_workers"`

	// **参数解释**： 最大演化轮次。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,100]。 **默认取值**： 不涉及
	SearchIterations int32 `json:"search_iterations"`

	// **参数解释**： 评估器返回值是越大越好还是越小越好。 **约束限制**： 不涉及 **取值范围**： 取值范围true，false。 **默认取值**： true
	SmallerBetter *bool `json:"smaller_better,omitempty"`

	// **参数解释**： 种群数量。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,32]。 **默认取值**： 不涉及
	SearchPopulationSize int32 `json:"search_population_size"`
}

func (o EvaluatorParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EvaluatorParameter struct{}"
	}

	return strings.Join([]string{"EvaluatorParameter", string(data)}, " ")
}
