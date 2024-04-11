package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoBatchOperateChildDto struct {

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`

	Params *BatchOperateChildDto `json:"params,omitempty"`
}

func (o RdmParamVoBatchOperateChildDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoBatchOperateChildDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoBatchOperateChildDto", string(data)}, " ")
}
