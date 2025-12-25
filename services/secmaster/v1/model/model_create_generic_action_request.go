package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGenericActionRequest Request Object
type CreateGenericActionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	Body *GenericActionRequestBody `json:"body,omitempty"`
}

func (o CreateGenericActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGenericActionRequest struct{}"
	}

	return strings.Join([]string{"CreateGenericActionRequest", string(data)}, " ")
}
