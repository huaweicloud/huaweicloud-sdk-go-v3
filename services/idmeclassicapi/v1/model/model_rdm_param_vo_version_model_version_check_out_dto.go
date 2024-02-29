package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelVersionCheckOutDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelVersionCheckOutDto `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelVersionCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelVersionCheckOutDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelVersionCheckOutDto", string(data)}, " ")
}
