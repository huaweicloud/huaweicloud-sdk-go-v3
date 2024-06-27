package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteChLtsConfigRequestBody 删除LTS配置
type DeleteChLtsConfigRequestBody struct {

	// LTS配置。
	LogConfigs []DeleteChLtsConfigRequestBodyLogConfigs `json:"log_configs"`
}

func (o DeleteChLtsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteChLtsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteChLtsConfigRequestBody", string(data)}, " ")
}
