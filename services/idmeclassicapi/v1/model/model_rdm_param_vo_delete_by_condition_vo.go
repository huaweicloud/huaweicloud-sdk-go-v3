package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoDeleteByConditionVo struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *DeleteByConditionVo `json:"params,omitempty"`
}

func (o RdmParamVoDeleteByConditionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoDeleteByConditionVo struct{}"
	}

	return strings.Join([]string{"RdmParamVoDeleteByConditionVo", string(data)}, " ")
}
