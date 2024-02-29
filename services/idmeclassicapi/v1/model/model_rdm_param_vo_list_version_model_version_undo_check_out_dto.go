package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModelVersionUndoCheckOutDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]VersionModelVersionUndoCheckOutDto `json:"params,omitempty"`
}

func (o RdmParamVoListVersionModelVersionUndoCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModelVersionUndoCheckOutDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModelVersionUndoCheckOutDto", string(data)}, " ")
}
