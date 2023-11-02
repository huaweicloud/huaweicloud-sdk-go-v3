package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePocketDetectionJobReq 创建靶点口袋发现请求体
type CreatePocketDetectionJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	Receptor *ReceptorDrugFile `json:"receptor"`

	Ligand *ProbeDrugFile `json:"ligand"`

	Params *PocketDetectionParamDto `json:"params,omitempty"`
}

func (o CreatePocketDetectionJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePocketDetectionJobReq struct{}"
	}

	return strings.Join([]string{"CreatePocketDetectionJobReq", string(data)}, " ")
}
