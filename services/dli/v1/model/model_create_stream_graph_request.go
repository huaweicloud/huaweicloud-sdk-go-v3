package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStreamGraphRequest Request Object
type CreateStreamGraphRequest struct {
	JobId string `json:"job_id"`

	Body *GenStreamGraphReq `json:"body,omitempty"`
}

func (o CreateStreamGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamGraphRequest struct{}"
	}

	return strings.Join([]string{"CreateStreamGraphRequest", string(data)}, " ")
}
