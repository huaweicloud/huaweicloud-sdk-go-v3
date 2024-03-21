package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableAsyncStatusLogRequest Request Object
type EnableAsyncStatusLogRequest struct {

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o EnableAsyncStatusLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAsyncStatusLogRequest struct{}"
	}

	return strings.Join([]string{"EnableAsyncStatusLogRequest", string(data)}, " ")
}
