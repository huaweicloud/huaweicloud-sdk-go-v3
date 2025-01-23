package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoUpdateDocRequestDto struct {
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *UpdateDocRequestDto `json:"params"`
}

func (o RdmParamVoUpdateDocRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoUpdateDocRequestDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoUpdateDocRequestDto", string(data)}, " ")
}
