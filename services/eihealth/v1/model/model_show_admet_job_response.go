package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdmetJobResponse Response Object
type ShowAdmetJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	MoleculeFile *MoleculeFileDto `json:"molecule_file,omitempty"`

	JobResult *JobResult `json:"job_result,omitempty"`

	// 作业结果信息
	PartFailedReason *[]FailedReasonRecord `json:"part_failed_reason,omitempty"`

	BaseModel *BaseModel `json:"base_model,omitempty"`

	// 模型信息
	Models *[]BasicDrugModel `json:"models,omitempty"`

	ClusterResult *ClusterJobRsp `json:"cluster_result,omitempty"`

	// 是否输出表征
	SaveFingerprint *bool `json:"save_fingerprint,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ShowAdmetJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdmetJobResponse struct{}"
	}

	return strings.Join([]string{"ShowAdmetJobResponse", string(data)}, " ")
}
