package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoMongPageRequest struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *MongPageRequest `json:"params,omitempty"`
}

func (o RdmParamVoMongPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoMongPageRequest struct{}"
	}

	return strings.Join([]string{"RdmParamVoMongPageRequest", string(data)}, " ")
}
