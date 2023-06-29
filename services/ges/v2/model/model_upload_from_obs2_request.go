package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFromObs2Request Request Object
type UploadFromObs2Request struct {
	Body *UploadFromObsReq `json:"body,omitempty"`
}

func (o UploadFromObs2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFromObs2Request struct{}"
	}

	return strings.Join([]string{"UploadFromObs2Request", string(data)}, " ")
}
