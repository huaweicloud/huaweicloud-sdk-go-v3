package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStartPositionResponse Response Object
type UpdateStartPositionResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateStartPositionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStartPositionResponse struct{}"
	}

	return strings.Join([]string{"UpdateStartPositionResponse", string(data)}, " ")
}
