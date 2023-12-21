package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistObjectIdsModifierDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *PersistObjectIdsModifierDto `json:"params,omitempty"`
}

func (o RdmParamVoPersistObjectIdsModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistObjectIdsModifierDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistObjectIdsModifierDto", string(data)}, " ")
}
