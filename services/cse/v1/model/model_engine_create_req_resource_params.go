package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineCreateReqResourceParams 微服务引擎资源参数
type EngineCreateReqResourceParams struct {

	// 是否自动刷新
	IsAutoRenew *int32 `json:"isAutoRenew,omitempty"`
}

func (o EngineCreateReqResourceParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineCreateReqResourceParams struct{}"
	}

	return strings.Join([]string{"EngineCreateReqResourceParams", string(data)}, " ")
}
