package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartPromptlyDataLevelTableCompareJobResponse Response Object
type StartPromptlyDataLevelTableCompareJobResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o StartPromptlyDataLevelTableCompareJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPromptlyDataLevelTableCompareJobResponse struct{}"
	}

	return strings.Join([]string{"StartPromptlyDataLevelTableCompareJobResponse", string(data)}, " ")
}
