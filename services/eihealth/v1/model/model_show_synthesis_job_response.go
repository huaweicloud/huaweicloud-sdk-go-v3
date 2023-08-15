package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSynthesisJobResponse Response Object
type ShowSynthesisJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	// 分子SMILES表达式
	Smiles *string `json:"smiles,omitempty"`

	Params         *SynthesisParamDto `json:"params,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowSynthesisJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSynthesisJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSynthesisJobResponse", string(data)}, " ")
}
