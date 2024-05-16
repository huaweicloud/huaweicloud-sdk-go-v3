package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeGraph2Response Response Object
type ResizeGraph2Response struct {

	// 变更图规格任务ID。请求失败时字段为空。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeGraph2Response struct{}"
	}

	return strings.Join([]string{"ResizeGraph2Response", string(data)}, " ")
}
