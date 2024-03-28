package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFileJobResourcesRequest Request Object
type UploadFileJobResourcesRequest struct {
	UserId *string `json:"USER-ID,omitempty"`

	Body *UploadResourcesRequestBody `json:"body,omitempty"`
}

func (o UploadFileJobResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFileJobResourcesRequest struct{}"
	}

	return strings.Join([]string{"UploadFileJobResourcesRequest", string(data)}, " ")
}
