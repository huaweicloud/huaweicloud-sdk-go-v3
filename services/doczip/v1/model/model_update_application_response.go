package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationResponse Response Object
type UpdateApplicationResponse struct {
	XTraceId       *string `json:"X-Trace-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationResponse", string(data)}, " ")
}
