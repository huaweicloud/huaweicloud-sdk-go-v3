package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConnectionResponse Response Object
type DeleteConnectionResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectionResponse struct{}"
	}

	return strings.Join([]string{"DeleteConnectionResponse", string(data)}, " ")
}
