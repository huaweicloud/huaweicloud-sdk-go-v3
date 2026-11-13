package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishApplicationResponse Response Object
type PublishApplicationResponse struct {
	XTraceId       *string `json:"X-Trace-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishApplicationResponse struct{}"
	}

	return strings.Join([]string{"PublishApplicationResponse", string(data)}, " ")
}
