package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateListenerRequest struct {

	// 监听器id
	ListenerId string `json:"listener_id" xml:"listener_id"`

	Body *UpdateListenerRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerRequest struct{}"
	}

	return strings.Join([]string{"UpdateListenerRequest", string(data)}, " ")
}
