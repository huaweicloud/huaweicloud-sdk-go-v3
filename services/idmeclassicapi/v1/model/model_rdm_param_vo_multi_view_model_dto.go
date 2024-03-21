package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoMultiViewModelDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *MultiViewModelCreateDto `json:"params,omitempty"`
}

func (o RdmParamVoMultiViewModelDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoMultiViewModelDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoMultiViewModelDto", string(data)}, " ")
}
