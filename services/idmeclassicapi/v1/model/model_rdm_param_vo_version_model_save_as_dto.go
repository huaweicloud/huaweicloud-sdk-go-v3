package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelSaveAsDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelSaveAsDto `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelSaveAsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelSaveAsDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelSaveAsDto", string(data)}, " ")
}
