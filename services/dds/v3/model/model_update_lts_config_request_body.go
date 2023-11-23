package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLtsConfigRequestBody struct {

	// 每个Item是实例需要设置的LTS策略。
	LtsConfigs []UpdateLtsConfigRequestBodyLtsConfigs `json:"lts_configs"`
}

func (o UpdateLtsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLtsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLtsConfigRequestBody", string(data)}, " ")
}
