package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteNoticeResponse Response Object
type BatchDeleteNoticeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteNoticeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteNoticeResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteNoticeResponse", string(data)}, " ")
}
