package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResizeServersRequest Request Object
type BatchResizeServersRequest struct {
	Body *BatchResizeServersReq `json:"body,omitempty"`
}

func (o BatchResizeServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResizeServersRequest struct{}"
	}

	return strings.Join([]string{"BatchResizeServersRequest", string(data)}, " ")
}
