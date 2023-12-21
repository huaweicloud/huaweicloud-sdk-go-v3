package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistObjectIdModifierDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *PersistObjectIdModifierDto `json:"params,omitempty"`
}

func (o RdmParamVoPersistObjectIdModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistObjectIdModifierDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistObjectIdModifierDto", string(data)}, " ")
}
