package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListPersistableModelSaveDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]PersistableModelSaveDto `json:"params,omitempty"`
}

func (o RdmParamVoListPersistableModelSaveDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListPersistableModelSaveDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListPersistableModelSaveDto", string(data)}, " ")
}
