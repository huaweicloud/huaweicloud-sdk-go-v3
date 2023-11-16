package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineConfigureReq 更新微服务引擎专享版配置请求体
type EngineConfigureReq struct {

	// 版本号
	Version string `json:"version"`
}

func (o EngineConfigureReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineConfigureReq struct{}"
	}

	return strings.Join([]string{"EngineConfigureReq", string(data)}, " ")
}
