package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoQueryRequestVo struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *QueryRequestVo `json:"params,omitempty"`
}

func (o RdmParamVoQueryRequestVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoQueryRequestVo struct{}"
	}

	return strings.Join([]string{"RdmParamVoQueryRequestVo", string(data)}, " ")
}
