package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBucketCredentialRequest Request Object
type CreateBucketCredentialRequest struct {
	Body *CreateBucketCredentialReq `json:"body,omitempty"`
}

func (o CreateBucketCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBucketCredentialRequest struct{}"
	}

	return strings.Join([]string{"CreateBucketCredentialRequest", string(data)}, " ")
}
