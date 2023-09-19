package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDigitalHumanBusinessCardResponse Response Object
type DeleteDigitalHumanBusinessCardResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDigitalHumanBusinessCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDigitalHumanBusinessCardResponse struct{}"
	}

	return strings.Join([]string{"DeleteDigitalHumanBusinessCardResponse", string(data)}, " ")
}
