package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSummaryBuildJobInfoResponse Response Object
type ShowSummaryBuildJobInfoResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *JobSummaryBodyResult `json:"result,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowSummaryBuildJobInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSummaryBuildJobInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSummaryBuildJobInfoResponse", string(data)}, " ")
}
