package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDigitalHumanBusinessCardResponse Response Object
type CreateDigitalHumanBusinessCardResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDigitalHumanBusinessCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDigitalHumanBusinessCardResponse struct{}"
	}

	return strings.Join([]string{"CreateDigitalHumanBusinessCardResponse", string(data)}, " ")
}
