package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistableModelSaveAsDto struct {
	Params *PersistableModelSaveAsDto `json:"params,omitempty"`

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`
}

func (o RdmParamVoPersistableModelSaveAsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistableModelSaveAsDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistableModelSaveAsDto", string(data)}, " ")
}
