package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelVersionUpdateAndCheckinDtoVersionModel struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelVersionUpdateAndCheckinDtoVersionModel `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelVersionUpdateAndCheckinDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelVersionUpdateAndCheckinDtoVersionModel struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelVersionUpdateAndCheckinDtoVersionModel", string(data)}, " ")
}
