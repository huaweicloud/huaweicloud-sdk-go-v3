package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModelVersionReviseDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]VersionModelVersionReviseDto `json:"params,omitempty"`
}

func (o RdmParamVoListVersionModelVersionReviseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModelVersionReviseDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModelVersionReviseDto", string(data)}, " ")
}
