package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartGraph2Response Response Object
type RestartGraph2Response struct {

	// 强制重启任务ID。请求失败时字段为空。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartGraph2Response struct{}"
	}

	return strings.Join([]string{"RestartGraph2Response", string(data)}, " ")
}
