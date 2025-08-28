package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePluginConfigReq 修改插件配置请求。
type UpdatePluginConfigReq struct {

	// 密钥。
	ApiKey *string `json:"api_key,omitempty"`
}

func (o UpdatePluginConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginConfigReq struct{}"
	}

	return strings.Join([]string{"UpdatePluginConfigReq", string(data)}, " ")
}
