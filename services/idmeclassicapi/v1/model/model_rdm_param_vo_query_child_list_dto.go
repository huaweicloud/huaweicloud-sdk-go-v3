package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoQueryChildListDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *QueryChildListDto `json:"params,omitempty"`
}

func (o RdmParamVoQueryChildListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoQueryChildListDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoQueryChildListDto", string(data)}, " ")
}
