package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFepJobResponse Response Object
type ShowFepJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	Receptor *ReceptorDrugFile `json:"receptor,omitempty"`

	// 是否加膜处理
	AddMembrane *bool `json:"add_membrane,omitempty"`

	// 配体列表
	Ligands *[]LigandPreviewDto `json:"ligands,omitempty"`

	Graph *FepGraphDto `json:"graph,omitempty"`

	Params *FepParamDto `json:"params,omitempty"`

	JobResult *JobResult `json:"job_result,omitempty"`

	// 部分失败原因和数量
	PartFailedReason *[]FailedReasonRecord `json:"part_failed_reason,omitempty"`
	HttpStatusCode   int                   `json:"-"`
}

func (o ShowFepJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFepJobResponse struct{}"
	}

	return strings.Join([]string{"ShowFepJobResponse", string(data)}, " ")
}
