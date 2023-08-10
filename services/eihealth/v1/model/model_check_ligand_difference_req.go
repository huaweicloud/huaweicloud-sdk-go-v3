package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckLigandDifferenceReq 配体差异计算请求体
type CheckLigandDifferenceReq struct {
	Method *CheckLigandDifferenceMethod `json:"method"`

	File *DrugFile `json:"file"`

	RefFile *DrugFile `json:"ref_file"`
}

func (o CheckLigandDifferenceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckLigandDifferenceReq struct{}"
	}

	return strings.Join([]string{"CheckLigandDifferenceReq", string(data)}, " ")
}
