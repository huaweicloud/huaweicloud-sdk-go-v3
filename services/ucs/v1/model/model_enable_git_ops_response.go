package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableGitOpsResponse Response Object
type EnableGitOpsResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableGitOpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableGitOpsResponse struct{}"
	}

	return strings.Join([]string{"EnableGitOpsResponse", string(data)}, " ")
}
