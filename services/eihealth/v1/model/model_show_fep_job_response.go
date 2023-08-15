package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFepJobResponse Response Object
type ShowFepJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	Receptor *ReceptorDrugFile `json:"receptor,omitempty"`

	// 配体列表
	Ligands *[]LigandPreviewDto `json:"ligands,omitempty"`

	Graph *FepGraphDto `json:"graph,omitempty"`

	Params *FepParamDto `json:"params,omitempty"`

	JobResult      *JobResult `json:"job_result,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowFepJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFepJobResponse struct{}"
	}

	return strings.Join([]string{"ShowFepJobResponse", string(data)}, " ")
}
