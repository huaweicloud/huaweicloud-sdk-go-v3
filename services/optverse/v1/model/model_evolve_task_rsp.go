package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EvolveTaskRsp struct {

	// **参数解释**： 任务标识符。 **约束限制**： 不涉及 **取值范围**： 长度[1-128] **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 优化任务名称。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,128]。 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// 优化任务描述
	Description *string `json:"description,omitempty"`

	// **参数解释**： 用户标识符。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,128]。 **默认取值**： 不涉及
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 优化结果存储路径。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,128]。 **默认取值**： 不涉及
	OutputPath *string `json:"output_path,omitempty"`

	// **参数解释**： 关联的算法设计项目。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,128]。 **默认取值**： 不涉及
	AlgorithmId *string `json:"algorithm_id,omitempty"`

	// **参数解释**： 关联的算法文件路径。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,256]。 **默认取值**： 不涉及
	AlgorithmFile *string `json:"algorithm_file,omitempty"`

	// **参数解释**： 算法函数名。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,256]。 **默认取值**： 不涉及
	AlgorithmFuncName *string `json:"algorithm_func_name,omitempty"`

	// **参数解释**： 算法进度。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,1]。 **默认取值**： 不涉及
	TaskProgress *float32 `json:"task_progress,omitempty"`

	// **参数解释**： 算法状态。 **约束限制**： 不涉及 **取值范围**： * DRAFT: 草稿 * PENDING: 初始化 * RUNNING: 运行中 * STOPPED: 已停止 * FINISHED: 已完成 * FAILED: 异常失败 **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： 评估器文件路径。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,65536]。 **默认取值**： 不涉及
	EvaluatorFile *string `json:"evaluator_file,omitempty"`

	// **参数解释**： 评估器算法函数名。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,256]。 **默认取值**： 不涉及
	EvaluatorFuncName *string `json:"evaluator_func_name,omitempty"`

	// **参数解释**： 评估基线文件路径。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,65536]。 **默认取值**： 不涉及
	EvaluatorBaseline *string `json:"evaluator_baseline,omitempty"`

	// **参数解释**： 评估基线算法函数名。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,256]。 **默认取值**： 不涉及
	EvaluatorBaselineFuncName *string `json:"evaluator_baseline_func_name,omitempty"`

	EvaluatorParameter *EvaluatorParameter `json:"evaluator_parameter,omitempty"`

	// **参数解释**： 关联CCE集群ID。 **约束限制**： 不涉及 **取值范围**： 取值范围[1,128]。 **默认取值**： 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 演化任务创建时间,单位毫秒。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,9999999999999]。 **默认取值**： 不涉及
	MetaCreateAt *int64 `json:"meta_create_at,omitempty"`

	// **参数解释**： 演化任务启动时间,单位毫秒。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,9999999999999]。 **默认取值**： 不涉及
	MetaStartAt *int64 `json:"meta_start_at,omitempty"`

	// **参数解释**： 演化任务完成时间,单位毫秒。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,9999999999999]。 **默认取值**： 不涉及
	MetaFinishAt *int64 `json:"meta_finish_at,omitempty"`

	// **参数解释**： 项目可见性。 **约束限制**： 不涉及 **取值范围**： * 公共: PUBLIC * 私有: PRIVATE **默认取值**： 不涉及
	Visibility *string `json:"visibility,omitempty"`
}

func (o EvolveTaskRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EvolveTaskRsp struct{}"
	}

	return strings.Join([]string{"EvolveTaskRsp", string(data)}, " ")
}
