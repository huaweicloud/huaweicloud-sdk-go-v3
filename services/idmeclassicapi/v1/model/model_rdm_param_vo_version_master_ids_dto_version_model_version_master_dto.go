package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionMasterIdsDtoVersionModelVersionMasterDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionMasterIdsDtoVersionModelVersionMasterDto `json:"params,omitempty"`
}

func (o RdmParamVoVersionMasterIdsDtoVersionModelVersionMasterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionMasterIdsDtoVersionModelVersionMasterDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionMasterIdsDtoVersionModelVersionMasterDto", string(data)}, " ")
}
