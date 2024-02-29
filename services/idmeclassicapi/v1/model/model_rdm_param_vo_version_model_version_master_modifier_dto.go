package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelVersionMasterModifierDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelVersionMasterModifierDto `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelVersionMasterModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelVersionMasterModifierDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelVersionMasterModifierDto", string(data)}, " ")
}
