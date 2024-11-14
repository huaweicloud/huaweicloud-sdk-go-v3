package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbCacheMappingResponse Response Object
type CreateDbCacheMappingResponse struct {

	// 创建内存加速任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDbCacheMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbCacheMappingResponse struct{}"
	}

	return strings.Join([]string{"CreateDbCacheMappingResponse", string(data)}, " ")
}
