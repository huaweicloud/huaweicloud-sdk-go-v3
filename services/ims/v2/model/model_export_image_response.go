package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportImageResponse Response Object
type ExportImageResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageResponse struct{}"
	}

	return strings.Join([]string{"ExportImageResponse", string(data)}, " ")
}
