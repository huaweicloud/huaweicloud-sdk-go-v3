package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelVersionMasterQueryDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelVersionMasterQueryDto `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelVersionMasterQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelVersionMasterQueryDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelVersionMasterQueryDto", string(data)}, " ")
}
