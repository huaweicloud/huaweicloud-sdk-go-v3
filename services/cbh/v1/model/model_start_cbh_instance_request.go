package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartCbhInstanceRequest struct {
	Body *StartCbhRequestBody `json:"body,omitempty"`
}

func (o StartCbhInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartCbhInstanceRequest struct{}"
	}

	return strings.Join([]string{"StartCbhInstanceRequest", string(data)}, " ")
}
