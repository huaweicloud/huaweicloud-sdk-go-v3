package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPocketDetectionJobResponse Response Object
type ShowPocketDetectionJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	Receptor *ReceptorDrugFile `json:"receptor,omitempty"`

	Ligand *ProbeDrugFile `json:"ligand,omitempty"`

	Params         *PocketDetectionParamDto `json:"params,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowPocketDetectionJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPocketDetectionJobResponse struct{}"
	}

	return strings.Join([]string{"ShowPocketDetectionJobResponse", string(data)}, " ")
}
