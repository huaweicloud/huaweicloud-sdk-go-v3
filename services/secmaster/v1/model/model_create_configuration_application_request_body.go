package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConfigurationApplicationRequestBody struct {

	// 配置列表
	Configuration []ConfigurationDetail `json:"configuration"`
}

func (o CreateConfigurationApplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationApplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConfigurationApplicationRequestBody", string(data)}, " ")
}
