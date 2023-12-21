package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelCreateDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelCreateDto `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelCreateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelCreateDto", string(data)}, " ")
}
