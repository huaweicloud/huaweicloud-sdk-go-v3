package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoQueryRequestCountVo struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *QueryRequestCountVo `json:"params,omitempty"`
}

func (o RdmParamVoQueryRequestCountVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoQueryRequestCountVo struct{}"
	}

	return strings.Join([]string{"RdmParamVoQueryRequestCountVo", string(data)}, " ")
}
