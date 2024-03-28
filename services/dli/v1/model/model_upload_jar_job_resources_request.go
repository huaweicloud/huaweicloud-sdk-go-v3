package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadJarJobResourcesRequest Request Object
type UploadJarJobResourcesRequest struct {
	UserId *string `json:"USER-ID,omitempty"`

	Body *UploadResourcesRequestBody `json:"body,omitempty"`
}

func (o UploadJarJobResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadJarJobResourcesRequest struct{}"
	}

	return strings.Join([]string{"UploadJarJobResourcesRequest", string(data)}, " ")
}
