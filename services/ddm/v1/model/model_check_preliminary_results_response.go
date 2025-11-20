package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPreliminaryResultsResponse Response Object
type CheckPreliminaryResultsResponse struct {

	// 关联的后端DN信息。
	PreCheckResults *[]PreCheckResult `json:"pre_check_results,omitempty"`

	// 工作流id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckPreliminaryResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPreliminaryResultsResponse struct{}"
	}

	return strings.Join([]string{"CheckPreliminaryResultsResponse", string(data)}, " ")
}
