package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncConnectionsNewRequestBody 同步连接请求体
type SyncConnectionsNewRequestBody struct {

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`
}

func (o SyncConnectionsNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncConnectionsNewRequestBody struct{}"
	}

	return strings.Join([]string{"SyncConnectionsNewRequestBody", string(data)}, " ")
}
