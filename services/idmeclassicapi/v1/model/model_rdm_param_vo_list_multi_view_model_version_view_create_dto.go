package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListMultiViewModelVersionViewCreateDto struct {

	// 参数对象。
	Params *[]MultiViewModelVersionViewCreateDto `json:"params,omitempty"`

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`
}

func (o RdmParamVoListMultiViewModelVersionViewCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListMultiViewModelVersionViewCreateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListMultiViewModelVersionViewCreateDto", string(data)}, " ")
}
