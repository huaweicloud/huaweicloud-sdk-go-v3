package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadPublicDataRequest Request Object
type DownloadPublicDataRequest struct {
	Body *DownloadPublicDataReq `json:"body,omitempty"`
}

func (o DownloadPublicDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadPublicDataRequest struct{}"
	}

	return strings.Join([]string{"DownloadPublicDataRequest", string(data)}, " ")
}
