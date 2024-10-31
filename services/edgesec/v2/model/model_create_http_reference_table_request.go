package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpReferenceTableRequest Request Object
type CreateHttpReferenceTableRequest struct {
	Body *CreateHttpReferenceTableRequestBody `json:"body,omitempty"`
}

func (o CreateHttpReferenceTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpReferenceTableRequest struct{}"
	}

	return strings.Join([]string{"CreateHttpReferenceTableRequest", string(data)}, " ")
}
