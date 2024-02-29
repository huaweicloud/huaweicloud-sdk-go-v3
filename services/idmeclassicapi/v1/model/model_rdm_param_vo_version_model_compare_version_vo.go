package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionModelCompareVersionVo struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *VersionModelCompareVersionVo `json:"params,omitempty"`
}

func (o RdmParamVoVersionModelCompareVersionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionModelCompareVersionVo struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionModelCompareVersionVo", string(data)}, " ")
}
