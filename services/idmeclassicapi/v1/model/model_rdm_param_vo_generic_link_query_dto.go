package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoGenericLinkQueryDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *GenericLinkQueryDto `json:"params,omitempty"`
}

func (o RdmParamVoGenericLinkQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoGenericLinkQueryDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoGenericLinkQueryDto", string(data)}, " ")
}
