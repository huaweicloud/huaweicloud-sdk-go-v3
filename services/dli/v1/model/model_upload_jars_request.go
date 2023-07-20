package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadJarsRequest Request Object
type UploadJarsRequest struct {
	UserId *string `json:"USER-ID,omitempty"`

	Body *UploadGroupPackageReq `json:"body,omitempty"`
}

func (o UploadJarsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadJarsRequest struct{}"
	}

	return strings.Join([]string{"UploadJarsRequest", string(data)}, " ")
}
