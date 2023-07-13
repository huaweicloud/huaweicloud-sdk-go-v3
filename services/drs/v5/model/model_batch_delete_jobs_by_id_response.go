package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobsByIdResponse Response Object
type BatchDeleteJobsByIdResponse struct {

	// 批量删除任务响应体。
	Jobs           *[]DeleteJobResp `json:"jobs,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchDeleteJobsByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsByIdResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsByIdResponse", string(data)}, " ")
}
