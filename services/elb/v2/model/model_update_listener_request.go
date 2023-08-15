package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateListenerRequest Request Object
type UpdateListenerRequest struct {

	// 监听器id
	ListenerId string `json:"listener_id"`

	Body *UpdateListenerRequestBody `json:"body,omitempty"`
}

func (o UpdateListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerRequest struct{}"
	}

	return strings.Join([]string{"UpdateListenerRequest", string(data)}, " ")
}
