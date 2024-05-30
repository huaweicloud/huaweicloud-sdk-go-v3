package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineConfigureReq 更新微服务引擎配置请求体。
type EngineConfigureReq struct {

	// authType安全认证类型，支持填写NONE和RBAC。
	AuthType string `json:"authType"`
}

func (o EngineConfigureReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineConfigureReq struct{}"
	}

	return strings.Join([]string{"EngineConfigureReq", string(data)}, " ")
}
