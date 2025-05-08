package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObjectRetrievalRequest Request Object
type CreateObjectRetrievalRequest struct {
	Body *CreateObjectRetrievalRequestBody `json:"body,omitempty"`
}

func (o CreateObjectRetrievalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectRetrievalRequest struct{}"
	}

	return strings.Join([]string{"CreateObjectRetrievalRequest", string(data)}, " ")
}
