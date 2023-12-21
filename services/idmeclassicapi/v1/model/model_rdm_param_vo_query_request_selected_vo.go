package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoQueryRequestSelectedVo struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *QueryRequestSelectedVo `json:"params,omitempty"`
}

func (o RdmParamVoQueryRequestSelectedVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoQueryRequestSelectedVo struct{}"
	}

	return strings.Join([]string{"RdmParamVoQueryRequestSelectedVo", string(data)}, " ")
}
