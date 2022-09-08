package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEventResponse struct {

	// 测试事件ID。
	Id *string `json:"id,omitempty"`

	// 测试事件名称。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEventResponse struct{}"
	}

	return strings.Join([]string{"UpdateEventResponse", string(data)}, " ")
}
