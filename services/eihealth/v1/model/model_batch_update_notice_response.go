package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateNoticeResponse Response Object
type BatchUpdateNoticeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateNoticeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNoticeResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateNoticeResponse", string(data)}, " ")
}
