package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelVersionCheckoutAndUpdateDtoVersionModel struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelVersionCheckoutAndUpdateDtoVersionModel `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelVersionCheckoutAndUpdateDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelVersionCheckoutAndUpdateDtoVersionModel struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelVersionCheckoutAndUpdateDtoVersionModel", string(data)}, " ")
}
