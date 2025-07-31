package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyImageInRegionInSafeModeResponse Response Object
type CopyImageInRegionInSafeModeResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyImageInRegionInSafeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageInRegionInSafeModeResponse struct{}"
	}

	return strings.Join([]string{"CopyImageInRegionInSafeModeResponse", string(data)}, " ")
}
