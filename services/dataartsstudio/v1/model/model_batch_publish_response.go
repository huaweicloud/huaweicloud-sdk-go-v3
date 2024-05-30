package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchPublishResponse Response Object
type BatchPublishResponse struct {
	Data           *BatchPublishResultData `json:"data,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchPublishResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPublishResponse struct{}"
	}

	return strings.Join([]string{"BatchPublishResponse", string(data)}, " ")
}
