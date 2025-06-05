package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDefaultBuildParametersParams struct {

	// 参数元数据
	Name *string `json:"name,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o ShowDefaultBuildParametersParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultBuildParametersParams struct{}"
	}

	return strings.Join([]string{"ShowDefaultBuildParametersParams", string(data)}, " ")
}
