package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistableModelCreateDto struct {
	Params *PersistableModelCreateDto `json:"params,omitempty"`

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`
}

func (o RdmParamVoPersistableModelCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistableModelCreateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistableModelCreateDto", string(data)}, " ")
}
