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

	JobResult      *JobResult `json:"job_result,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowDockingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDockingJobResponse struct{}"
	}

	return strings.Join([]string{"ShowDockingJobResponse", string(data)}, " ")
}
