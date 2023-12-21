package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListPersistableModelCreateDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]PersistableModelCreateDto `json:"params,omitempty"`
}

func (o RdmParamVoListPersistableModelCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListPersistableModelCreateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListPersistableModelCreateDto", string(data)}, " ")
}
