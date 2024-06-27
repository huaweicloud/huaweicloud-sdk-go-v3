package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChLtsConfigRequestBody 打开LTS配置
type CreateChLtsConfigRequestBody struct {

	// 日志配置信息。
	LogConfigs []CreateChLtsConfigRequestBodyLogConfigs `json:"log_configs"`
}

func (o CreateChLtsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChLtsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CreateChLtsConfigRequestBody", string(data)}, " ")
}
