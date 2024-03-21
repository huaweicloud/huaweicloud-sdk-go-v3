package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistObjectIdDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *PersistObjectIdDto `json:"params,omitempty"`
}

func (o RdmParamVoPersistObjectIdDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistObjectIdDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistObjectIdDto", string(data)}, " ")
}
