package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTargetOptJobResponse Response Object
type ShowTargetOptJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	Receptor *TargetOptReceptor `json:"receptor,omitempty"`

	Ligand *TargetOptLigand `json:"ligand,omitempty"`

	MdParams       *MdParam `json:"md_params,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowTargetOptJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTargetOptJobResponse struct{}"
	}

	return strings.Join([]string{"ShowTargetOptJobResponse", string(data)}, " ")
}
