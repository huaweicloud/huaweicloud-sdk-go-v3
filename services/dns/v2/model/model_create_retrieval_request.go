package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRetrievalRequest Request Object
type CreateRetrievalRequest struct {
	Body *CreateRetrievalRequestBody `json:"body,omitempty"`
}

func (o CreateRetrievalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetrievalRequest struct{}"
	}

	return strings.Join([]string{"CreateRetrievalRequest", string(data)}, " ")
}
