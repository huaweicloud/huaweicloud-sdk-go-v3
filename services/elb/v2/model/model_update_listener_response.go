package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateListenerResponse Response Object
type UpdateListenerResponse struct {
	Listener       *ListenerResp `json:"listener,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerResponse struct{}"
	}

	return strings.Join([]string{"UpdateListenerResponse", string(data)}, " ")
}
