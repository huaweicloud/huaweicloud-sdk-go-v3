package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMessageReceiveConfigResponse Response Object
type UpdateMessageReceiveConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMessageReceiveConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageReceiveConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateMessageReceiveConfigResponse", string(data)}, " ")
}
