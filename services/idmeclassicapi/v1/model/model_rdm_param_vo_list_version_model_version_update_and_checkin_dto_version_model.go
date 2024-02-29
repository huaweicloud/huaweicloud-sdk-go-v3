package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModelVersionUpdateAndCheckinDtoVersionModel struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]VersionModelVersionUpdateAndCheckinDtoVersionModel `json:"params,omitempty"`
}

func (o RdmParamVoListVersionModelVersionUpdateAndCheckinDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModelVersionUpdateAndCheckinDtoVersionModel struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModelVersionUpdateAndCheckinDtoVersionModel", string(data)}, " ")
}
