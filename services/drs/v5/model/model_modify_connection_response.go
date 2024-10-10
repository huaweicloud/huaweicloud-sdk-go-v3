package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyConnectionResponse Response Object
type ModifyConnectionResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ModifyConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyConnectionResponse struct{}"
	}

	return strings.Join([]string{"ModifyConnectionResponse", string(data)}, " ")
}
