package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateListenerRequestBody struct {
	Listener *UpdateListenerReq `json:"listener"`
}

func (o UpdateListenerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateListenerRequestBody", string(data)}, " ")
}
