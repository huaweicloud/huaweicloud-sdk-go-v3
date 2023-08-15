package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandGraph2Response Response Object
type ExpandGraph2Response struct {

	// 扩副本任务ID。请求失败时字段为空。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandGraph2Response struct{}"
	}

	return strings.Join([]string{"ExpandGraph2Response", string(data)}, " ")
}
