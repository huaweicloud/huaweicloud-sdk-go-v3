package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyImageCrossRegionResponse Response Object
type CopyImageCrossRegionResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyImageCrossRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageCrossRegionResponse struct{}"
	}

	return strings.Join([]string{"CopyImageCrossRegionResponse", string(data)}, " ")
}
