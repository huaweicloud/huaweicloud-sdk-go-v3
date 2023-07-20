package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFilesRequest Request Object
type UploadFilesRequest struct {
	UserId *string `json:"USER-ID,omitempty"`

	Body *UploadGroupPackageReq `json:"body,omitempty"`
}

func (o UploadFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFilesRequest struct{}"
	}

	return strings.Join([]string{"UploadFilesRequest", string(data)}, " ")
}
