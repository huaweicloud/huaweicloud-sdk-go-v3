package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableGitOpsResponse Response Object
type DisableGitOpsResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableGitOpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableGitOpsResponse struct{}"
	}

	return strings.Join([]string{"DisableGitOpsResponse", string(data)}, " ")
}
