package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogstashConfResponse Response Object
type DeleteLogstashConfResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteLogstashConfResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogstashConfResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogstashConfResponse", string(data)}, " ")
}
