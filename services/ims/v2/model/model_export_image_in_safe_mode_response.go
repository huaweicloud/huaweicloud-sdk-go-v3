package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportImageInSafeModeResponse Response Object
type ExportImageInSafeModeResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportImageInSafeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageInSafeModeResponse struct{}"
	}

	return strings.Join([]string{"ExportImageInSafeModeResponse", string(data)}, " ")
}
