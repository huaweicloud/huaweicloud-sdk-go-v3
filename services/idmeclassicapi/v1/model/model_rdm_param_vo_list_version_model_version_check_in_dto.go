package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModelVersionCheckInDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]VersionModelVersionCheckInDto `json:"params,omitempty"`
}

func (o RdmParamVoListVersionModelVersionCheckInDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModelVersionCheckInDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModelVersionCheckInDto", string(data)}, " ")
}
