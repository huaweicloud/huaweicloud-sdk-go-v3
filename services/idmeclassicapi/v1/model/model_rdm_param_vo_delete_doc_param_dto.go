package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoDeleteDocParamDto struct {
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *PersistObjectIdsDto `json:"params,omitempty"`
}

func (o RdmParamVoDeleteDocParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoDeleteDocParamDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoDeleteDocParamDto", string(data)}, " ")
}
