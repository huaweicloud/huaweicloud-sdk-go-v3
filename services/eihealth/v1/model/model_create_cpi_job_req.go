package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCpiJobReq 创建CPI作业请求体
type CreateCpiJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	// 受体文件列表
	Receptors []CpiReceptor `json:"receptors"`

	// 小分子
	Ligands []MoleculeFileDto `json:"ligands"`
}

func (o CreateCpiJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCpiJobReq struct{}"
	}

	return strings.Join([]string{"CreateCpiJobReq", string(data)}, " ")
}
