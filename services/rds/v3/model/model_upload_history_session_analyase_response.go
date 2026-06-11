package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadHistorySessionAnalyaseResponse Response Object
type UploadHistorySessionAnalyaseResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadHistorySessionAnalyaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadHistorySessionAnalyaseResponse struct{}"
	}

	return strings.Join([]string{"UploadHistorySessionAnalyaseResponse", string(data)}, " ")
}
