package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitAsyncJobResponse Response Object
type CommitAsyncJobResponse struct {
	Job            *AsyncCommitJobResp `json:"job,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CommitAsyncJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitAsyncJobResponse struct{}"
	}

	return strings.Join([]string{"CommitAsyncJobResponse", string(data)}, " ")
}
