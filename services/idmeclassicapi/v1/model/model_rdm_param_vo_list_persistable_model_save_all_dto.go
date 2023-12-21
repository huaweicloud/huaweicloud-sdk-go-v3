package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListPersistableModelSaveAllDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]PersistableModelSaveAllDto `json:"params,omitempty"`
}

func (o RdmParamVoListPersistableModelSaveAllDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListPersistableModelSaveAllDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListPersistableModelSaveAllDto", string(data)}, " ")
}
