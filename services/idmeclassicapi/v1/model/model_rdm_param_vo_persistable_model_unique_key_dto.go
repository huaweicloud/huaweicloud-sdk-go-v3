package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistableModelUniqueKeyDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *PersistableModelUniqueKeyDto `json:"params,omitempty"`
}

func (o RdmParamVoPersistableModelUniqueKeyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistableModelUniqueKeyDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistableModelUniqueKeyDto", string(data)}, " ")
}
