package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartInstanceRequest Request Object
type StartInstanceRequest struct {
	Body *CommonCbhRequestBody `json:"body,omitempty"`
}

func (o StartInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceRequest", string(data)}, " ")
}
