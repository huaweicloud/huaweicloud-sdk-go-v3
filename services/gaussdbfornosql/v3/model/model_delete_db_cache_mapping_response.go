package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDbCacheMappingResponse Response Object
type DeleteDbCacheMappingResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDbCacheMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbCacheMappingResponse struct{}"
	}

	return strings.Join([]string{"DeleteDbCacheMappingResponse", string(data)}, " ")
}
