package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartGraph2Response Response Object
type StartGraph2Response struct {

	// 启动图任务ID。请求失败时字段为空。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartGraph2Response struct{}"
	}

	return strings.Join([]string{"StartGraph2Response", string(data)}, " ")
}
