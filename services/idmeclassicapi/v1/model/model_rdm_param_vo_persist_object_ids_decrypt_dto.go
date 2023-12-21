package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistObjectIdsDecryptDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *PersistObjectIdsDecryptDto `json:"params,omitempty"`
}

func (o RdmParamVoPersistObjectIdsDecryptDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistObjectIdsDecryptDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistObjectIdsDecryptDto", string(data)}, " ")
}
