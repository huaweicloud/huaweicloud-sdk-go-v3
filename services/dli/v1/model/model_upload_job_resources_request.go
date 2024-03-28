package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadJobResourcesRequest Request Object
type UploadJobResourcesRequest struct {
	UserId *string `json:"USER-ID,omitempty"`

	Body *UploadJobResourcesRequestBody `json:"body,omitempty"`
}

func (o UploadJobResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadJobResourcesRequest struct{}"
	}

	return strings.Join([]string{"UploadJobResourcesRequest", string(data)}, " ")
}
