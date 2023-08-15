package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineUpdateReq 升级微服务引擎专享版请求体
type EngineUpdateReq struct {

	// 版本号
	Version string `json:"version"`
}

func (o EngineUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineUpdateReq struct{}"
	}

	return strings.Join([]string{"EngineUpdateReq", string(data)}, " ")
}
