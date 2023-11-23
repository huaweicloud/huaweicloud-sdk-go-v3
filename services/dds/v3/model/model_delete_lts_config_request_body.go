package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteLtsConfigRequestBody struct {

	// 需要解除的LTS配置列表，一个实例解除多种日志配置需要填写多个item。
	LtsConfigs []DeleteLtsConfigRequestBodyLtsConfigs `json:"lts_configs"`
}

func (o DeleteLtsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLtsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteLtsConfigRequestBody", string(data)}, " ")
}
