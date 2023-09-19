package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDigitalHumanBusinessCardResponse Response Object
type UpdateDigitalHumanBusinessCardResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDigitalHumanBusinessCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDigitalHumanBusinessCardResponse struct{}"
	}

	return strings.Join([]string{"UpdateDigitalHumanBusinessCardResponse", string(data)}, " ")
}
