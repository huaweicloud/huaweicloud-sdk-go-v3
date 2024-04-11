package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoCompareVersionVo struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *CompareVersionVo `json:"params,omitempty"`
}

func (o RdmParamVoCompareVersionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoCompareVersionVo struct{}"
	}

	return strings.Join([]string{"RdmParamVoCompareVersionVo", string(data)}, " ")
}
