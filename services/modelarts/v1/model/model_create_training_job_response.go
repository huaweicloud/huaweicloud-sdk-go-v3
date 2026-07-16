package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainingJobResponse Response Object
type CreateTrainingJobResponse struct {

	// **参数解释**：训练作业类型。 **取值范围**： - job：普通作业 - federated_pool_job：资源池联邦作业 - edge_job：边缘作业 - hetero_job：异构作业 - mrs_job：MRS作业 - autosearch_job：自动化搜索作业 - diag_job：诊断作业 - visualization_job：可视化作业
	Kind *string `json:"kind,omitempty"`

	Metadata *JobMetadataResponse `json:"metadata,omitempty"`

	Status *Status `json:"status,omitempty"`

	Algorithm *JobAlgorithmResponse `json:"algorithm,omitempty"`

	// **参数解释**：异构训练作业的任务列表。
	Tasks *[]TaskResponse `json:"tasks,omitempty"`

	Spec *SpecResponse `json:"spec,omitempty"`

	Endpoints *JobEndpointsResp `json:"endpoints,omitempty"`

	FtjobConfig    *MasJobConfig `json:"ftjob_config,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateTrainingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingJobResponse struct{}"
	}

	return strings.Join([]string{"CreateTrainingJobResponse", string(data)}, " ")
}
