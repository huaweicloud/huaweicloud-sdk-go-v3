package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelVersionReviseDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelVersionReviseDto `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelVersionReviseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelVersionReviseDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelVersionReviseDto", string(data)}, " ")
}
