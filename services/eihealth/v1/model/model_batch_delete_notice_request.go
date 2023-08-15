package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteNoticeRequest Request Object
type BatchDeleteNoticeRequest struct {
	Body *BatchDeleteNoticeReq `json:"body,omitempty"`
}

func (o BatchDeleteNoticeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteNoticeRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteNoticeRequest", string(data)}, " ")
}
