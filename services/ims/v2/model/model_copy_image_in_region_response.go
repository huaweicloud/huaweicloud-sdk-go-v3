package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyImageInRegionResponse Response Object
type CopyImageInRegionResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyImageInRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageInRegionResponse struct{}"
	}

	return strings.Join([]string{"CopyImageInRegionResponse", string(data)}, " ")
}
