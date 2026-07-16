package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDevServersActionRequest Request Object
type BatchDevServersActionRequest struct {
	Body *DevServerBatchRequest `json:"body,omitempty"`
}

func (o BatchDevServersActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDevServersActionRequest struct{}"
	}

	return strings.Join([]string{"BatchDevServersActionRequest", string(data)}, " ")
}
