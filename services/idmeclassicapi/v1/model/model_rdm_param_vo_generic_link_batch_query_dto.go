package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoGenericLinkBatchQueryDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *GenericLinkBatchQueryDto `json:"params,omitempty"`
}

func (o RdmParamVoGenericLinkBatchQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoGenericLinkBatchQueryDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoGenericLinkBatchQueryDto", string(data)}, " ")
}
