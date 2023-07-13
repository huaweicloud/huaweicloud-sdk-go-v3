package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDownloadResourceStatDataRequest Request Object
type BatchDownloadResourceStatDataRequest struct {
	Body *BatchQueryStatReq `json:"body,omitempty"`
}

func (o BatchDownloadResourceStatDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDownloadResourceStatDataRequest struct{}"
	}

	return strings.Join([]string{"BatchDownloadResourceStatDataRequest", string(data)}, " ")
}
