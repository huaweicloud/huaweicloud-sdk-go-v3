package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublicationResponse Response Object
type CreatePublicationResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePublicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicationResponse struct{}"
	}

	return strings.Join([]string{"CreatePublicationResponse", string(data)}, " ")
}
