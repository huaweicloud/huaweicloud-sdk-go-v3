package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDockJobReq 创建分子对接作业请求体
type CreateDockJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	// 受体文件列表
	Receptors []DockingReceptorDto `json:"receptors"`

	// 配体文件列表，当前仅支持1个
	Ligands []LigandDto `json:"ligands"`
}

func (o CreateDockJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDockJobReq struct{}"
	}

	return strings.Join([]string{"CreateDockJobReq", string(data)}, " ")
}
