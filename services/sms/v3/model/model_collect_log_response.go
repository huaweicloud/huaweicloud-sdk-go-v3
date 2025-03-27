package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectLogResponse Response Object
type CollectLogResponse struct {

	// 上传迁移任务的日志成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CollectLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectLogResponse struct{}"
	}

	return strings.Join([]string{"CollectLogResponse", string(data)}, " ")
}
