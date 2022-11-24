package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEventResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventResponse struct{}"
	}

	return strings.Join([]string{"DeleteEventResponse", string(data)}, " ")
}
