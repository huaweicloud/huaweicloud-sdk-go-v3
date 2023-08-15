package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventResponse Response Object
type CreateEventResponse struct {

	// 测试事件ID。
	Id *string `json:"id,omitempty"`

	// 测试事件名称。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventResponse struct{}"
	}

	return strings.Join([]string{"CreateEventResponse", string(data)}, " ")
}
