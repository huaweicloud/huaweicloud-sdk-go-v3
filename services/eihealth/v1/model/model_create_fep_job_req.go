package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建自由能微扰作业请求体
type CreateFepJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	Receptor *DrugFile `json:"receptor"`

	// 配体列表
	Ligands []LigandPreviewDto `json:"ligands"`

	Graph *FepGraphDto `json:"graph"`

	Params *FepParamDto `json:"params"`
}

func (o CreateFepJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFepJobReq struct{}"
	}

	return strings.Join([]string{"CreateFepJobReq", string(data)}, " ")
}
