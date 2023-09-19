package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Create2DDigitalHumanVideoResponse Response Object
type Create2DDigitalHumanVideoResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o Create2DDigitalHumanVideoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Create2DDigitalHumanVideoResponse struct{}"
	}

	return strings.Join([]string{"Create2DDigitalHumanVideoResponse", string(data)}, " ")
}
