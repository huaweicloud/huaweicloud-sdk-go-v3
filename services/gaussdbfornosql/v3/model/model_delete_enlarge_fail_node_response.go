package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnlargeFailNodeResponse Response Object
type DeleteEnlargeFailNodeResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEnlargeFailNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnlargeFailNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteEnlargeFailNodeResponse", string(data)}, " ")
}
