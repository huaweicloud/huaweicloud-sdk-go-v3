package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAppRequest Request Object
type UploadAppRequest struct {
	Body *UploadAppReq `json:"body,omitempty"`
}

func (o UploadAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAppRequest struct{}"
	}

	return strings.Join([]string{"UploadAppRequest", string(data)}, " ")
}
