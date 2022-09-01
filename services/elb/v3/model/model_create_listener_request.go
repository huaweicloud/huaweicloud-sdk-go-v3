package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateListenerRequest struct {
	Body *CreateListenerRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerRequest struct{}"
	}

	return strings.Join([]string{"CreateListenerRequest", string(data)}, " ")
}
