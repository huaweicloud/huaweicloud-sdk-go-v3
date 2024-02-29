package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelVersionUndoCheckOutDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelVersionUndoCheckOutDto `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelVersionUndoCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelVersionUndoCheckOutDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelVersionUndoCheckOutDto", string(data)}, " ")
}
