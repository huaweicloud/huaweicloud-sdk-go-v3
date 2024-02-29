package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModel struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	// 参数对象。
	Params *[]VersionModel `json:"params,omitempty"`
}

func (o RdmParamVoListVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModel struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModel", string(data)}, " ")
}
