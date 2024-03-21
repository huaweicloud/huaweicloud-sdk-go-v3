package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoStatisticsPvo struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *StatisticsPvo `json:"params,omitempty"`
}

func (o RdmParamVoStatisticsPvo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoStatisticsPvo struct{}"
	}

	return strings.Join([]string{"RdmParamVoStatisticsPvo", string(data)}, " ")
}
