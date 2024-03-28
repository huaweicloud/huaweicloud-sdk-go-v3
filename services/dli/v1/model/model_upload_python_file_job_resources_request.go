package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadPythonFileJobResourcesRequest Request Object
type UploadPythonFileJobResourcesRequest struct {
	UserId *string `json:"USER-ID,omitempty"`

	Body *UploadResourcesRequestBody `json:"body,omitempty"`
}

func (o UploadPythonFileJobResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadPythonFileJobResourcesRequest struct{}"
	}

	return strings.Join([]string{"UploadPythonFileJobResourcesRequest", string(data)}, " ")
}
