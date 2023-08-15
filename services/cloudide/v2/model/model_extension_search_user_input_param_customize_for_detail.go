package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionSearchUserInputParamCustomizeForDetail struct {

	// 插件ID列表
	Ids string `json:"ids"`
}

func (o ExtensionSearchUserInputParamCustomizeForDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionSearchUserInputParamCustomizeForDetail struct{}"
	}

	return strings.Join([]string{"ExtensionSearchUserInputParamCustomizeForDetail", string(data)}, " ")
}
