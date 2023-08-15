package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEventSourceResponse Response Object
type DeleteEventSourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEventSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventSourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteEventSourceResponse", string(data)}, " ")
}
