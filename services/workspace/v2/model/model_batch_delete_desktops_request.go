package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopsRequest Request Object
type BatchDeleteDesktopsRequest struct {
	Body *DeleteDesktopsReq `json:"body,omitempty"`
}

func (o BatchDeleteDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopsRequest", string(data)}, " ")
}
