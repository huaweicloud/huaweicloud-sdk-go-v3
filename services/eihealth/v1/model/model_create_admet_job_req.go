package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAdmetJobReq 创建分子属性预测作业请求体
type CreateAdmetJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	MoleculeFile *MoleculeFileDto `json:"molecule_file"`

	// 基模型id
	BaseModelId *string `json:"base_model_id,omitempty"`

	// 模型id列表
	ModelIds *[]string `json:"model_ids,omitempty"`

	// 是否输出表征，仅专业版平台支持
	SaveFingerprint *bool `json:"save_fingerprint,omitempty"`
}

func (o CreateAdmetJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAdmetJobReq struct{}"
	}

	return strings.Join([]string{"CreateAdmetJobReq", string(data)}, " ")
}
