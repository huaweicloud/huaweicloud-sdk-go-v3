package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoGenericLinkTypeModifierDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *GenericLinkTypeModifierDto `json:"params,omitempty"`
}

func (o RdmParamVoGenericLinkTypeModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoGenericLinkTypeModifierDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoGenericLinkTypeModifierDto", string(data)}, " ")
}
