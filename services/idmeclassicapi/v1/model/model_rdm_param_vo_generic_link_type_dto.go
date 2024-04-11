package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoGenericLinkTypeDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *GenericLinkTypeDto `json:"params,omitempty"`
}

func (o RdmParamVoGenericLinkTypeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoGenericLinkTypeDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoGenericLinkTypeDto", string(data)}, " ")
}
