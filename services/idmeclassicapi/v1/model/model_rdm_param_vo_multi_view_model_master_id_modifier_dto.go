package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoMultiViewModelMasterIdModifierDto struct {
	Params *MultiViewModelMasterIdModifierDto `json:"params,omitempty"`

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`
}

func (o RdmParamVoMultiViewModelMasterIdModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoMultiViewModelMasterIdModifierDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoMultiViewModelMasterIdModifierDto", string(data)}, " ")
}
