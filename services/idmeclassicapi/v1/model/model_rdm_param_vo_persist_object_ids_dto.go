package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistObjectIdsDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *PersistObjectIdsDto `json:"params,omitempty"`
}

func (o RdmParamVoPersistObjectIdsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistObjectIdsDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistObjectIdsDto", string(data)}, " ")
}
