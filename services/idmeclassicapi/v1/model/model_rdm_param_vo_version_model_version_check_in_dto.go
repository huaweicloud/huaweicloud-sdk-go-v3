package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelVersionCheckInDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelVersionCheckInDto `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelVersionCheckInDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelVersionCheckInDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelVersionCheckInDto", string(data)}, " ")
}
