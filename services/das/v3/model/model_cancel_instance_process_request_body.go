package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelInstanceProcessRequestBody 删除会话请求体
type CancelInstanceProcessRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 会话列表
	Processes []CancelInstanceProcessInfo `json:"processes"`
}

func (o CancelInstanceProcessRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelInstanceProcessRequestBody struct{}"
	}

	return strings.Join([]string{"CancelInstanceProcessRequestBody", string(data)}, " ")
}
