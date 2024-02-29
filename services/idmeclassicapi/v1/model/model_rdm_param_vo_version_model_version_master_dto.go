package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelVersionMasterDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelVersionMasterDto `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelVersionMasterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelVersionMasterDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelVersionMasterDto", string(data)}, " ")
}
