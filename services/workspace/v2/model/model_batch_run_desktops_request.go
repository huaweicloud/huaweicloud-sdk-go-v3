package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRunDesktopsRequest Request Object
type BatchRunDesktopsRequest struct {
	Body *BatchActionDesktopsReq `json:"body,omitempty"`
}

func (o BatchRunDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRunDesktopsRequest struct{}"
	}

	return strings.Join([]string{"BatchRunDesktopsRequest", string(data)}, " ")
}
