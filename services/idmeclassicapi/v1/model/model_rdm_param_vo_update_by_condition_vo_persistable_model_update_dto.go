package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoUpdateByConditionVoPersistableModelUpdateDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *UpdateByConditionVoPersistableModelUpdateDto `json:"params,omitempty"`
}

func (o RdmParamVoUpdateByConditionVoPersistableModelUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoUpdateByConditionVoPersistableModelUpdateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoUpdateByConditionVoPersistableModelUpdateDto", string(data)}, " ")
}
