package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateSpaceAnalysisTaskResponse struct {
	// 执行时间，毫秒为单位的时间戳

	ExecutionTime  *int64 `json:"execution_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateSpaceAnalysisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSpaceAnalysisTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateSpaceAnalysisTaskResponse", string(data)}, " ")
}
