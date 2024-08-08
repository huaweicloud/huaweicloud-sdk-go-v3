package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDockingJobResponse Response Object
type ShowDockingJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	// 受体文件列表
	Receptors *[]DockingReceptorDto `json:"receptors,omitempty"`

	// 配体文件列表，当前仅支持1个
	Ligands *[]LigandDto `json:"ligands,omitempty"`

	// 引擎
	Engine *string `json:"engine,omitempty"`

	JobResult *JobResult `json:"job_result,omitempty"`

	// 部分失败原因和数量
	PartFailedReason *[]FailedReasonRecord `json:"part_failed_reason,omitempty"`

	ClusterResult  *ClusterJobRsp `json:"cluster_result,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowDockingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDockingJobResponse struct{}"
	}

	return strings.Join([]string{"ShowDockingJobResponse", string(data)}, " ")
}
