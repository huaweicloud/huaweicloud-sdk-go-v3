package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModelVersionUpdateDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]VersionModelVersionUpdateDto `json:"params,omitempty"`
}

func (o RdmParamVoListVersionModelVersionUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModelVersionUpdateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModelVersionUpdateDto", string(data)}, " ")
}
