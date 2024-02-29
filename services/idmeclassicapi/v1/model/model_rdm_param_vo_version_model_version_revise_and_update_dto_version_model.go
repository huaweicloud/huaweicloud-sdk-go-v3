package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelVersionReviseAndUpdateDtoVersionModel struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelVersionReviseAndUpdateDtoVersionModel `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelVersionReviseAndUpdateDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelVersionReviseAndUpdateDtoVersionModel struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelVersionReviseAndUpdateDtoVersionModel", string(data)}, " ")
}
