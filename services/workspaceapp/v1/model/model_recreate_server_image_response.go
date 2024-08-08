package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecreateServerImageResponse Response Object
type RecreateServerImageResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecreateServerImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecreateServerImageResponse struct{}"
	}

	return strings.Join([]string{"RecreateServerImageResponse", string(data)}, " ")
}
