package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListPersistableModelUpdateDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]PersistableModelUpdateDto `json:"params,omitempty"`
}

func (o RdmParamVoListPersistableModelUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListPersistableModelUpdateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListPersistableModelUpdateDto", string(data)}, " ")
}
