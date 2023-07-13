package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopGraph2Response Response Object
type StopGraph2Response struct {

	// 关闭图任务ID。请求失败时为空。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopGraph2Response struct{}"
	}

	return strings.Join([]string{"StopGraph2Response", string(data)}, " ")
}
