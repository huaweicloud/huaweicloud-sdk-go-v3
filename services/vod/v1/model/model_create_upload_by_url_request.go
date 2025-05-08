package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUploadByUrlRequest Request Object
type CreateUploadByUrlRequest struct {
	Body *UploadByUrlReq `json:"body,omitempty"`
}

func (o CreateUploadByUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUploadByUrlRequest struct{}"
	}

	return strings.Join([]string{"CreateUploadByUrlRequest", string(data)}, " ")
}
