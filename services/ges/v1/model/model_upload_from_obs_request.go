package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UploadFromObsRequest struct {
	Body *UploadFromObsReq `json:"body,omitempty" xml:"body"`
}

func (o UploadFromObsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFromObsRequest struct{}"
	}

	return strings.Join([]string{"UploadFromObsRequest", string(data)}, " ")
}
