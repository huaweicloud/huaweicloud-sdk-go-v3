package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateNoticeRequest Request Object
type BatchUpdateNoticeRequest struct {
	Body *BatchUpdateNoticeReq `json:"body,omitempty"`
}

func (o BatchUpdateNoticeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNoticeRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateNoticeRequest", string(data)}, " ")
}
