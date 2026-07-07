package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAnalysisSessionStatusResponse Response Object
type ShowAnalysisSessionStatusResponse struct {

	// 任务状态: pending/running/finished/error/timeout
	Status *string `json:"status,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAnalysisSessionStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnalysisSessionStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAnalysisSessionStatusResponse", string(data)}, " ")
}
