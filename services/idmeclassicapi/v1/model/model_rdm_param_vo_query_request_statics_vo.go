package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoQueryRequestStaticsVo struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *QueryRequestStaticsVo `json:"params,omitempty"`
}

func (o RdmParamVoQueryRequestStaticsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoQueryRequestStaticsVo struct{}"
	}

	return strings.Join([]string{"RdmParamVoQueryRequestStaticsVo", string(data)}, " ")
}
