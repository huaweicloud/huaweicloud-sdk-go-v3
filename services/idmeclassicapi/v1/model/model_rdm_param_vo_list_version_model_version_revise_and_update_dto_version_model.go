package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModelVersionReviseAndUpdateDtoVersionModel struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]VersionModelVersionReviseAndUpdateDtoVersionModel `json:"params,omitempty"`
}

func (o RdmParamVoListVersionModelVersionReviseAndUpdateDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModelVersionReviseAndUpdateDtoVersionModel struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModelVersionReviseAndUpdateDtoVersionModel", string(data)}, " ")
}
