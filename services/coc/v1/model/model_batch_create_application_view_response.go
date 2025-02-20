package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateApplicationViewResponse Response Object
type BatchCreateApplicationViewResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchCreateApplicationViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewResponse", string(data)}, " ")
}
