package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateListenerRequest struct {
	Body *CreateListenerRequestBody `json:"body,omitempty"`
}

func (o CreateListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerRequest struct{}"
	}

	return strings.Join([]string{"CreateListenerRequest", string(data)}, " ")
}
