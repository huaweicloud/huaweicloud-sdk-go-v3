package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartInstanceResponse Response Object
type StartInstanceResponse struct {

	// 启动数据库的异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceResponse", string(data)}, " ")
}
