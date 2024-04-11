package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoQueryParentDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *QueryParentDto `json:"params,omitempty"`
}

func (o RdmParamVoQueryParentDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoQueryParentDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoQueryParentDto", string(data)}, " ")
}
