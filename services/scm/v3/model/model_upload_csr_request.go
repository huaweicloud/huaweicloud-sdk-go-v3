package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadCsrRequest Request Object
type UploadCsrRequest struct {
	Body *UploadCsrRequestBody `json:"body,omitempty"`
}

func (o UploadCsrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadCsrRequest struct{}"
	}

	return strings.Join([]string{"UploadCsrRequest", string(data)}, " ")
}
