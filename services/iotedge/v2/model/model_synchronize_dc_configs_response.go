package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SynchronizeDcConfigsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SynchronizeDcConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynchronizeDcConfigsResponse struct{}"
	}

	return strings.Join([]string{"SynchronizeDcConfigsResponse", string(data)}, " ")
}
