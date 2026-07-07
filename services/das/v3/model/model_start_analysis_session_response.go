package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAnalysisSessionResponse Response Object
type StartAnalysisSessionResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartAnalysisSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAnalysisSessionResponse struct{}"
	}

	return strings.Join([]string{"StartAnalysisSessionResponse", string(data)}, " ")
}
