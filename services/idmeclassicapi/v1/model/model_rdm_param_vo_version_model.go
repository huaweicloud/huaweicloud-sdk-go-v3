package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModel struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModel `json:"params,omitempty"`
}

func (o RdmParamVoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModel struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModel", string(data)}, " ")
}
