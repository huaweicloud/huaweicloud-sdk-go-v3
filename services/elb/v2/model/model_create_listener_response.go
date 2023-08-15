package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateListenerResponse Response Object
type CreateListenerResponse struct {
	Listener       *ListenerResp `json:"listener,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerResponse struct{}"
	}

	return strings.Join([]string{"CreateListenerResponse", string(data)}, " ")
}
