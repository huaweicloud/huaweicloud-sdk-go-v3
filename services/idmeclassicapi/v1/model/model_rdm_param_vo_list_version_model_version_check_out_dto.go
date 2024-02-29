package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModelVersionCheckOutDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]VersionModelVersionCheckOutDto `json:"params,omitempty"`
}

func (o RdmParamVoListVersionModelVersionCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModelVersionCheckOutDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModelVersionCheckOutDto", string(data)}, " ")
}
