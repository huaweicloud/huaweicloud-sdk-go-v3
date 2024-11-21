package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMolDockingJobReq 单分子预对接请求体。
type CreateMolDockingJobReq struct {
	Receptor *ReceptorDto `json:"receptor"`

	Ligand *DrugFile `json:"ligand"`

	// 引擎，默认为AUTODOCK_VINA。
	Engine *string `json:"engine,omitempty"`
}

func (o CreateMolDockingJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMolDockingJobReq struct{}"
	}

	return strings.Join([]string{"CreateMolDockingJobReq", string(data)}, " ")
}
