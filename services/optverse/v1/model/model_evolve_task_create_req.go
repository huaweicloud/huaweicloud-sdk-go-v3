package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EvolveTaskCreateReq struct {

	// **参数解释**： 优化任务名称。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,64]。 **默认取值**： 不涉及
	Name string `json:"name"`

	// **参数解释**： 优化任务描述。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,65536]。 **默认取值**： 不涉及
	Description string `json:"description"`

	// **参数解释**： 优化结果存储路径。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,256]。 **默认取值**： 不涉及
	OutputPath string `json:"output_path"`

	// **参数解释**： 关联的算法设计项目。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,128]。 **默认取值**： 不涉及
	AlgorithmId string `json:"algorithm_id"`

	// **参数解释**： 关联的算法文件路径。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,256]。 **默认取值**： 不涉及
	AlgorithmFile *string `json:"algorithm_file,omitempty"`

	// **参数解释**： 算法函数名。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,256]。 **默认取值**： 不涉及
	AlgorithmFuncName *string `json:"algorithm_func_name,omitempty"`

	// **参数解释**： 评估器文件路径。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,65536]。 **默认取值**： 不涉及
	EvaluatorFile string `json:"evaluator_file"`

	// **参数解释**： 评估器算法函数名。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,256]。 **默认取值**： 不涉及
	EvaluatorFuncName string `json:"evaluator_func_name"`

	// **参数解释**： 评估基线文件路径。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,65536]。 **默认取值**： 不涉及
	EvaluatorBaseline string `json:"evaluator_baseline"`

	// **参数解释**： 评估基线算法函数名。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,256]。 **默认取值**： 不涉及
	EvaluatorBaselineFuncName string `json:"evaluator_baseline_func_name"`

	EvaluatorParameter *EvaluatorParameter `json:"evaluator_parameter"`
}

func (o EvolveTaskCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EvolveTaskCreateReq struct{}"
	}

	return strings.Join([]string{"EvolveTaskCreateReq", string(data)}, " ")
}
