package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResizeServersResponse Response Object
type BatchResizeServersResponse struct {
	OrderId *string `json:"order_id,omitempty"`

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchResizeServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResizeServersResponse struct{}"
	}

	return strings.Join([]string{"BatchResizeServersResponse", string(data)}, " ")
}
