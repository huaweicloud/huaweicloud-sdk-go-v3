package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDefaultBuildParametersResult struct {

	// 参数定义
	Name *string `json:"name,omitempty"`

	// 参数元数据列表
	Params *[]ShowDefaultBuildParametersParams `json:"params,omitempty"`
}

func (o ShowDefaultBuildParametersResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultBuildParametersResult struct{}"
	}

	return strings.Join([]string{"ShowDefaultBuildParametersResult", string(data)}, " ")
}
