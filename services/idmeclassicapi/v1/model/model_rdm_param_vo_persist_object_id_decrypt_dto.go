package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistObjectIdDecryptDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *PersistObjectIdDecryptDto `json:"params,omitempty"`
}

func (o RdmParamVoPersistObjectIdDecryptDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistObjectIdDecryptDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistObjectIdDecryptDto", string(data)}, " ")
}
