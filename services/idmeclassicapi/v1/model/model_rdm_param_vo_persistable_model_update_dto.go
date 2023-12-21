package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistableModelUpdateDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *PersistableModelUpdateDto `json:"params,omitempty"`
}

func (o RdmParamVoPersistableModelUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistableModelUpdateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistableModelUpdateDto", string(data)}, " ")
}
