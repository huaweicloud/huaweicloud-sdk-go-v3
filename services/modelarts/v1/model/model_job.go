package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Job 创建训练作业的作业请求体和相应体。
type Job struct {

	// **参数解释**：训练作业类型。 **约束限制**：不涉及。 **取值范围**： - job：普通作业 - federated_pool_job：资源池联邦作业 - edge_job：边缘作业 - hetero_job：异构作业 - mrs_job：MRS作业 - autosearch_job：自动化搜索作业 - diag_job：诊断作业 - visualization_job：可视化作业  **默认取值**：job。
	Kind string `json:"kind"`

	Metadata *JobMetadata `json:"metadata"`

	Algorithm *JobAlgorithm `json:"algorithm,omitempty"`

	// **参数解释**：任务列表。该功能暂未实现。 **约束限制**：不涉及。
	Tasks *[]Task `json:"tasks,omitempty"`

	Spec *Spec `json:"spec,omitempty"`

	Endpoints *JobEndpointsReq `json:"endpoints,omitempty"`

	// **参数解释**：类型。 **约束限制**：不涉及。 **取值范围**：SFT（全量微调）、PRETRAIN（预训练）、LORA（lora微调）、DPO（dpo强化学习）、RFT（rft强化学习）
	TrainType *string `json:"train_type,omitempty"`

	FtjobConfig *MasJobConfig `json:"ftjob_config,omitempty"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
