package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoMultiViewModelVersionViewCreateDto struct {
	Params *MultiViewModelVersionViewCreateDto `json:"params,omitempty"`

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`
}

func (o RdmParamVoMultiViewModelVersionViewCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoMultiViewModelVersionViewCreateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoMultiViewModelVersionViewCreateDto", string(data)}, " ")
}
