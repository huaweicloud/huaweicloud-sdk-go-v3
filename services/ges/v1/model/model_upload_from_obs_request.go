package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFromObsRequest Request Object
type UploadFromObsRequest struct {
	Body *UploadFromObsReq `json:"body,omitempty"`
}

func (o UploadFromObsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFromObsRequest struct{}"
	}

	return strings.Join([]string{"UploadFromObsRequest", string(data)}, " ")
}
