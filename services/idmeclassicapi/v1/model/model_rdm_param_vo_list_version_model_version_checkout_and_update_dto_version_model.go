package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModelVersionCheckoutAndUpdateDtoVersionModel struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]VersionModelVersionCheckoutAndUpdateDtoVersionModel `json:"params,omitempty"`
}

func (o RdmParamVoListVersionModelVersionCheckoutAndUpdateDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModelVersionCheckoutAndUpdateDtoVersionModel struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModelVersionCheckoutAndUpdateDtoVersionModel", string(data)}, " ")
}
