package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCpiJobResponse Response Object
type ShowCpiJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	// 受体文件列表
	Receptors *[]ReceptorDrugFile `json:"receptors,omitempty"`

	// 小分子
	Ligands        *[]MoleculeFileDto `json:"ligands,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowCpiJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCpiJobResponse struct{}"
	}

	return strings.Join([]string{"ShowCpiJobResponse", string(data)}, " ")
}
