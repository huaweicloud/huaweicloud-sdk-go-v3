package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageServerResponse Response Object
type CreateImageServerResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateImageServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageServerResponse struct{}"
	}

	return strings.Join([]string{"CreateImageServerResponse", string(data)}, " ")
}
