package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTargetOptJobReq 创建靶点优化作业请求体
type CreateTargetOptJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	Receptor *TargetOptReceptor `json:"receptor"`

	Ligand *TargetOptLigand `json:"ligand,omitempty"`

	MdParams *MdParam `json:"md_params,omitempty"`
}

func (o CreateTargetOptJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTargetOptJobReq struct{}"
	}

	return strings.Join([]string{"CreateTargetOptJobReq", string(data)}, " ")
}
