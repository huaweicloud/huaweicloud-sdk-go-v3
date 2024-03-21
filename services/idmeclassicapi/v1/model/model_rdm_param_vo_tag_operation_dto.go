package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoTagOperationDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *TagOperationDto `json:"params,omitempty"`
}

func (o RdmParamVoTagOperationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoTagOperationDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoTagOperationDto", string(data)}, " ")
}
