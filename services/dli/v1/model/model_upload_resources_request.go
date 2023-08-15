package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadResourcesRequest Request Object
type UploadResourcesRequest struct {
	UserId *string `json:"USER-ID,omitempty"`

	Body *UploadPackageGroupReq `json:"body,omitempty"`
}

func (o UploadResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadResourcesRequest struct{}"
	}

	return strings.Join([]string{"UploadResourcesRequest", string(data)}, " ")
}
